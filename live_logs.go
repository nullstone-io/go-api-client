package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type LiveLogs struct {
	Client *Client
}

func (l LiveLogs) path(stackId int64, runUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/runs/%s/live_logs", l.Client.Config.OrgName, stackId, runUid)
}

func (l LiveLogs) Watch(ctx context.Context, stackId int64, runUid uuid.UUID) (<-chan types.LiveLogMessage, error) {
	c, err := l.connect(stackId, runUid)
	if err != nil {
		return nil, err
	}

	// Live log messages are sent to this channel
	ch := make(chan types.LiveLogMessage)

	// done signals when websocket ReadMessage fails - usually due to a closed connection
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer close(ch)
		for {
			_, data, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
					// The stream may not be ready yet, let's wait a second and retry
					select {
						case <-ctx.Done():
							return
						case <-time.After(time.Second):
					}
					if c, err = l.connect(stackId, runUid); err != nil {
						log.Println("error retrying connection: %w", err)
						return
					}
					continue
				}
				return
			}
			var msg types.LiveLogMessage
			if err := json.Unmarshal(data, &msg); err != nil {
				fmt.Println("error reading live log message: %w", err)
				continue
			}
			ch <- msg
		}
	}()

	go func() {
		// This goroutine blocks until either:
		// 1. The server closed the connection causing the done channel closed
		// 2. ctx was cancelled (likely due to Ctrl+C)
		// The websocket connection is forcefully closed just in case
		defer c.Close()
		select {
		case <-done:
			return
		case <-ctx.Done():
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("error closing websocket: %w", err)
				return
			}
		}
	}()

	return ch, nil
}

func (l LiveLogs) connect(stackId int64, runUid uuid.UUID) (*websocket.Conn, error) {
	endpoint, err := url.Parse(l.Client.Config.BaseAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}
	endpoint.Path = l.path(stackId, runUid)
	endpoint.Scheme = strings.Replace(endpoint.Scheme, "http", "ws", 1)
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+l.Client.Config.ApiKey)
	c, res, err := websocket.DefaultDialer.Dial(endpoint.String(), headers)
	if err != nil {
		return nil, fmt.Errorf("error connecting to live logs: %w", err)
	}
	if err := response.Verify(res); err != nil {
		return nil, fmt.Errorf("error connecting to live logs: %w", err)
	}
	return c, nil
}
