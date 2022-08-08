package live_logs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"strings"
	"time"
)

const (
	EndOfTransmission = "\x04"
)

func NewReconnectingStreamer(ctx context.Context, endpoint string, headers http.Header) (*ReconnectingStreamer, error) {
	s := &ReconnectingStreamer{
		Endpoint: endpoint,
		Headers:  headers,
		done:     make(chan struct{}),
	}
	if err := s.connect(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

type ReconnectingStreamer struct {
	Endpoint string
	Headers  http.Header

	conn *websocket.Conn

	// done signals when websocket ReadMessage fails - usually due to a closed connection
	done chan struct{}
}

func (s *ReconnectingStreamer) Stream(ctx context.Context) chan types.LiveLogMessage {
	// This channel is returned and represents the stream of messages from the websocket endpoint
	ch := make(chan types.LiveLogMessage)
	go s.readLoop(ctx, ch)
	go s.watchClose(ctx)
	return ch
}

func (s *ReconnectingStreamer) connect(ctx context.Context) error {
	c, res, err := websocket.DefaultDialer.DialContext(ctx, s.Endpoint, s.Headers)
	if err != nil {
		return err
	}
	s.conn = c
	return response.Verify(res)
}

// readLoop runs until closed by context cancellation or close error from the server
func (s *ReconnectingStreamer) readLoop(ctx context.Context, ch chan types.LiveLogMessage) {
	defer close(s.done)
	defer close(ch)
	for {
		_, data, err := s.conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				// The stream may not be ready yet, let's wait a second and retry
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Second):
				}
				if err := s.connect(ctx); err != nil {
					ch <- types.LiveLogMessage{
						Source:  "error",
						Content: fmt.Sprintf("error retrying connection: %s\n", err),
					}
					return
				}
				continue
			}
			return
		}
		var msg types.LiveLogMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			ch <- types.LiveLogMessage{
				Source:  "error",
				Content: fmt.Sprintf("error reading live log message: %s\n", err),
			}
			continue
		}
		eot := strings.HasSuffix(msg.Content, EndOfTransmission)
		if eot {
			msg.Content = strings.TrimSuffix(msg.Content, EndOfTransmission)
		}
		if msg.Content != "" {
			ch <- msg
		}
		if eot {
			// End of transmission characters terminates the stream
			return
		}
	}
}

func (s *ReconnectingStreamer) watchClose(ctx context.Context) {
	// The websocket connection is forcefully closed just in case
	defer s.close()

	// This blocks until either:
	// 1. The server closed the connection causing the done channel closed
	// 2. ctx was cancelled (likely due to Ctrl+C)
	select {
	case <-s.done:
		return
	case <-ctx.Done():
		err := s.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			return
		}
	}
}

func (s *ReconnectingStreamer) close() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}
