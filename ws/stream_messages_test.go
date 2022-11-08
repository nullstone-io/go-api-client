package ws

import (
	"context"
	"fmt"
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

func TestStreamMessages(t *testing.T) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1204,
		WriteBufferSize: 1204,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	tests := []struct {
		name             string
		messages         []types.Message
		serverEmitsClose bool
		clientTimeout    time.Duration
	}{
		{
			name: "send 3 messages and EOT",
			messages: []types.Message{
				{
					Context: types.DeployPhaseInit,
					Content: "message 1\n",
				},
				{
					Context: types.DeployPhaseCheckout,
					Content: "message 2\n",
				},
				{
					Context: types.DeployPhaseCheckout,
					Content: "message 3\n",
				},
			},
			serverEmitsClose: true,
		},
		{
			name: "send 1 message, client cancels",
			messages: []types.Message{
				{
					Context: types.DeployPhaseInit,
					Content: "message 1\n",
				},
			},
			clientTimeout: 100 * time.Millisecond,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/endpoint" {
					conn, err := upgrader.Upgrade(w, r, nil)
					if err != nil {
						http.Error(w, fmt.Sprintf("error upgrading websocket: %s", err), http.StatusInternalServerError)
						return
					}
					assert.NotNil(t, conn, "websocket connection")

					for _, msg := range test.messages {
						conn.WriteJSON(msg)
					}
					if test.serverEmitsClose {
						closeData := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "end of transmission")
						conn.WriteMessage(websocket.CloseMessage, closeData)
					}

					msgType, data, err := conn.ReadMessage()
					log.Println(msgType, data, err)
				}
			}))

			ctx := context.Background()
			if test.clientTimeout > 0 {
				// This simulates a user cancelling the stream from the client side
				// This verifies that the
				var cancel func()
				ctx, cancel = context.WithTimeout(ctx, test.clientTimeout)
				defer cancel()
			}

			u, err := url.Parse(server.URL)
			require.NoError(t, err, "parse server url")
			u.Path = "/endpoint"
			u.Scheme = strings.Replace(u.Scheme, "http", "ws", 1)
			ch := StreamMessages(ctx, u.String(), http.Header{}, RetryInfinite(time.Millisecond))
			assert.NotNil(t, ch, "stream channel")
			got := make([]types.Message, 0)
			for msg := range ch {
				got = append(got, msg)
			}
			assert.Equal(t, test.messages, got)
		})
	}
}
