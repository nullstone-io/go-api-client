package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

// StreamObject performs a websocket request against an endpoint
// This blocks until the initial "hydrate" response containing the full object as it currently exists
// The returned channel will emit update objects as the server streams updates
func StreamObject[T any, TUpdate any](ctx context.Context, endpoint string, headers http.Header, retryFn StreamerRetryFunc) (*T, <-chan types.StreamObject[TUpdate], error) {
	s := Streamer{
		Endpoint: endpoint,
		Headers:  headers,
		RetryFn:  retryFn,
	}
	rawMsgs := s.Stream(ctx)

	var initial T
	if msg, err := readStreamMessage(rawMsgs); err != nil {
		return nil, nil, err
	} else if msg == nil {
		return nil, nil, nil
	} else if msg.Context != "hydrate" {
		return nil, nil, fmt.Errorf("did not receive initial object from stream")
	} else if err := json.Unmarshal([]byte(msg.Content), &initial); err != nil {
		return nil, nil, fmt.Errorf("error reading initial object from stream: %w", err)
	}

	ch := make(chan types.StreamObject[TUpdate])
	go func() {
		defer close(ch)
		for {
			if msg, err := readStreamMessage(rawMsgs); err != nil {
				ch <- types.StreamObject[TUpdate]{Err: err}
			} else if msg == nil {
				return
			} else {
				var so types.StreamObject[TUpdate]
				so.Err = json.Unmarshal([]byte(msg.Content), &so.Object)
				ch <- so
			}
		}
	}()

	return &initial, ch, nil
}

func readStreamMessage(ch <-chan []byte) (*types.Message, error) {
	raw, ok := <-ch
	if !ok {
		return nil, nil
	}
	var msg types.Message
	if err := json.Unmarshal(raw, &msg); err != nil {
		return nil, fmt.Errorf("error reading message from stream: %w", err)
	}
	return &msg, nil
}
