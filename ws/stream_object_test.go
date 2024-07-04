package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestStreamObject(t *testing.T) {
	type hydrateObject struct {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	type updateObject struct {
		Id     int64  `json:"id"`
		Status string `json:"status"`
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1204,
		WriteBufferSize: 1204,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	tests := []struct {
		name                         string
		messages                     []types.Message
		closeServerAfterTestMessages bool
		clientTimeout                time.Duration
		wantInitial                  hydrateObject
		wantUpdates                  []types.StreamObject[updateObject]
	}{
		{
			name: "send hydrate and 2 updates",
			messages: []types.Message{
				{
					Context: "hydrate",
					Content: `{"id":1,"name":"test1","status":"calculating"}`,
				},
				{
					Context: "updated",
					Content: `{"id":1,"status":"running"}`,
				},
				{
					Context: "updated",
					Content: `{"id":1,"status":"completed"}`,
				},
			},
			closeServerAfterTestMessages: true,
			wantInitial: hydrateObject{
				Id:     1,
				Name:   "test1",
				Status: "calculating",
			},
			wantUpdates: []types.StreamObject[updateObject]{
				{
					Object: updateObject{
						Id:     1,
						Status: "running",
					},
				},
				{
					Object: updateObject{
						Id:     1,
						Status: "completed",
					},
				},
			},
		},
		{
			name: "send hydrate, client cancels",
			messages: []types.Message{
				{
					Context: "hydrate",
					Content: `{"id":1,"name":"test1","status":"calculating"}`,
				},
			},
			clientTimeout: 10 * time.Millisecond,
			wantInitial: hydrateObject{
				Id:     1,
				Name:   "test1",
				Status: "calculating",
			},
			wantUpdates: []types.StreamObject[updateObject]{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
			defer cancel()
			if test.clientTimeout > 0 {
				// This simulates a user cancelling the stream from the client side
				// This verifies that the
				var cancel func()
				ctx, cancel = context.WithTimeout(ctx, test.clientTimeout)
				defer cancel()
			}

			router := mux.NewRouter()
			router.Methods(http.MethodGet).
				Path("/endpoint").
				HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					conn, err := upgrader.Upgrade(w, r, nil)
					if err != nil {
						http.Error(w, fmt.Sprintf("error upgrading websocket: %s", err), http.StatusInternalServerError)
						return
					}
					require.NotNil(t, conn, "websocket connection")
					defer conn.Close()

					for _, msg := range test.messages {
						conn.WriteJSON(msg)
					}
					if test.closeServerAfterTestMessages {
						closeData := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "end of transmission")
						conn.WriteMessage(websocket.CloseMessage, closeData)
					}

					msgType, data, err := conn.ReadMessage()
					log.Println(msgType, data, err)
				})
			server := httptest.NewServer(router)

			u, err := url.Parse(server.URL)
			require.NoError(t, err, "parse server url")
			u.Path = "/endpoint"
			u.Scheme = strings.Replace(u.Scheme, "http", "ws", 1)
			gotInitial, ch, err := StreamObject[hydrateObject, updateObject](ctx, u.String(), http.Header{}, RetryInfinite(time.Millisecond))
			require.NoError(t, err, "unexpected error")
			require.NotNil(t, ch, "stream channel")
			if assert.NotNil(t, gotInitial, "hydrate object not nil") {
				assert.Equal(t, test.wantInitial, *gotInitial, "hydrate object")
			}
			gotUpdates := make([]types.StreamObject[updateObject], 0)
			for update := range ch {
				gotUpdates = append(gotUpdates, update)
			}
			assert.Equal(t, test.wantUpdates, gotUpdates, "updates")
		})
	}
}
