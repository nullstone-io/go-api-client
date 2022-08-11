package live_logs

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"gopkg.in/nullstone-io/go-api-client.v0/websocket"
	"net/http"
)

func Stream(ctx context.Context, endpoint string, headers http.Header, retryFn websocket.StreamerRetryFunc) <-chan types.LiveLogMessage {
	s := websocket.Streamer{
		Endpoint: endpoint,
		Headers:  headers,
		RetryFn:  retryFn,
	}
	rawMsgs := s.Stream(ctx)
	logMsgs := make(chan types.LiveLogMessage)
	go func() {
		defer close(logMsgs)
		for raw := range rawMsgs {
			var msg types.LiveLogMessage
			if err := json.Unmarshal(raw, &msg); err != nil {
				msg = types.LiveLogMessage{
					Source:  "error",
					Content: fmt.Sprintf("error reading log: %s\n", err),
				}
			}
			logMsgs <- msg
		}
	}()
	return logMsgs
}
