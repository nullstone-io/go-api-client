package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type StreamerRetryFunc func(ctx context.Context, readErr error) (retry bool, delay time.Duration)

type Streamer struct {
	Endpoint string
	Headers  http.Header
	RetryFn  StreamerRetryFunc

	conn *websocket.Conn
}

func (s *Streamer) Stream(ctx context.Context) <-chan []byte {
	ch := make(chan []byte)
	done := make(chan struct{})
	go s.streamLoop(ctx, ch, done)
	go s.watchCancel(ctx, done)
	return ch
}

func (s *Streamer) streamLoop(ctx context.Context, ch chan []byte, done chan struct{}) {
	defer close(done)
	defer close(ch)

	connected := false

	for {
		if !connected {
			if err := s.connect(ctx); err != nil {
				if s.shouldReconnect(ctx, err) {
					// TODO: How do we tell the user that we're reconnecting?
					// shouldReconnect pauses, let's restart the loop
					continue
				}
				return
			}
			connected = true
		}

		_, data, err := s.conn.ReadMessage()
		if err != nil {
			connected = false
			// This error could result from transport issues, corrupt data, or a closed message sent from the server
			// The connection must be re-established, or we terminate the streamer
			if s.shouldReconnect(ctx, err) {
				// TODO: How do we tell the user that we're reconnecting?
				continue
			}
			return
		}
		ch <- data
	}
}

func (s *Streamer) watchCancel(ctx context.Context, done chan struct{}) {
	// This blocks until either:
	// 1. The server closed the connection causing the done channel closed
	// 2. ctx was cancelled (likely due to Ctrl+C)
	select {
	case <-done:
		// The connection closed, we're done
		return
	case <-ctx.Done():
		// The context was cancelled, let's tell the server to close
		err := s.conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "cancelled"))
		if err != nil {
			// If we're unable to write a close message, the connection is likely closed
			return
		}
	}
}

// connect establishes a new websocket connection
// This will only fail if the server is down, cannot handle the request, or fails to upgrade to websocket connection
// This does not fail if the server sends a close message and does not need to handle auto
func (s *Streamer) connect(ctx context.Context) error {
	var err error
	s.conn, _, err = websocket.DefaultDialer.DialContext(ctx, s.Endpoint, s.Headers)
	return err
}

func (s *Streamer) shouldReconnect(ctx context.Context, err error) bool {
	// Let's use the RetryFn to determine whether we should attempt to reconnect
	// If RetryFn is nil, we terminate the streamer
	if s.RetryFn == nil {
		return false
	}
	retry, delay := s.RetryFn(ctx, err)
	if !retry {
		return false
	}
	if delay <= 0 {
		delay = time.Microsecond
	}
	select {
	case <-ctx.Done():
		// If we cancel during the wait, let's not attempt to reconnect
		return false
	case <-time.After(delay):
	}

	return true
}
