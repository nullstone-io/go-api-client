package websocket

import (
	"context"
	"time"
)

func RetryInfinite(delay time.Duration) StreamerRetryFunc {
	return func(ctx context.Context, readErr error) (bool, time.Duration) {
		return true, delay
	}
}
