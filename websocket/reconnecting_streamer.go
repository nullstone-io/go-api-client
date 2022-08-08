package websocket

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"log"
	"net/http"
	"os"
	"time"
)

func NewReconnectingStreamer[T any](ctx context.Context, endpoint string, headers http.Header) (*ReconnectingStreamer[T], error) {
	s := &ReconnectingStreamer[T]{
		Endpoint: endpoint,
		Headers:  headers,
		done:     make(chan struct{}),
		logger:   log.New(os.Stderr, "", 0),
	}
	if err := s.connect(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

type ReconnectingStreamer[T any] struct {
	Endpoint string
	Headers  http.Header

	conn   *websocket.Conn
	logger *log.Logger

	// done signals when websocket ReadMessage fails - usually due to a closed connection
	done chan struct{}
}

func (s *ReconnectingStreamer[T]) Stream(ctx context.Context) chan T {
	// This channel is returned and represents the stream of messages from the websocket endpoint
	ch := make(chan T)
	go s.readLoop(ctx, ch)
	go s.watchClose(ctx)
	return ch
}

func (s *ReconnectingStreamer[T]) connect(ctx context.Context) error {
	c, res, err := websocket.DefaultDialer.DialContext(ctx, s.Endpoint, s.Headers)
	if err != nil {
		return err
	}
	s.conn = c
	return response.Verify(res)
}

// readLoop runs until closed by context cancellation or close error from the server
func (s *ReconnectingStreamer[T]) readLoop(ctx context.Context, ch chan T) {
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
					s.logger.Printf("error retrying connection: %s\n", err)
					return
				}
				continue
			}
			return
		}
		var msg T
		if err := json.Unmarshal(data, &msg); err != nil {
			s.logger.Printf("error reading live log message: %s\n", err)
			continue
		}
		ch <- msg
	}
}

func (s *ReconnectingStreamer[T]) watchClose(ctx context.Context) {
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
			s.logger.Printf("error closing websocket: %s\n", err)
			return
		}
	}
}

func (s *ReconnectingStreamer[T]) close() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}
