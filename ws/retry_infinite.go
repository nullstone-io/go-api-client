package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"time"
)

func RetryInfinite(delay time.Duration) StreamerRetryFunc {
	return func(ctx context.Context, readErr error) (bool, time.Duration) {
		if websocket.IsCloseError(readErr, websocket.CloseNormalClosure) {
			return false, 0
		}
		return true, delay
	}
}
