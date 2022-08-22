package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

func StreamLogs(ctx context.Context, endpoint string, headers http.Header, retryFn StreamerRetryFunc) <-chan types.Message {
	s := Streamer{
		Endpoint: endpoint,
		Headers:  headers,
		RetryFn:  retryFn,
	}
	rawMsgs := s.Stream(ctx)
	logMsgs := make(chan types.Message)
	go func() {
		defer close(logMsgs)
		for raw := range rawMsgs {
			var msg types.Message
			if err := json.Unmarshal(raw, &msg); err != nil {
				msg = types.Message{
					Source:  "error",
					Content: fmt.Sprintf("error reading log: %s\n", err),
				}
			}
			logMsgs <- msg
		}
	}()
	return logMsgs
}
