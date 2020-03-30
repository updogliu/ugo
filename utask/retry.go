package utask

import (
	"context"
	"time"

	"github.com/updogliu/ugo/ulog"
	"github.com/updogliu/ugo/utime"
)

// Returns nil on the first time `f()` returns nil, even if by that time `ctx` has
// been cancelled. Otherwise returns the last error returned by `f()`. If `f()` has
// never got a chance to run, returns `ctx.Err()`.
func Retry(ctx context.Context, taskName string, retryGapMs int64, f func() error) error {
	// Suppress logging more than one error within 5 sec.
	logCooldown := utime.NewUnreadyCooldown(5 * time.Second)

	ticker := time.NewTicker(retryGapMs * time.Millisecond)
	defer ticker.Stop()

	var lastErr error
	for {
		select {
		case <-ctx.Done():
			if lastErr == nil {
				return ctx.Err()
			}
			return lastErr

		default:
			lastErr = f()
			if lastErr == nil {
				return nil
			}

			if logCooldown.Ready() {
				ulog.Errorf("Still retrying %s. Last error: %v", taskName, lastErr)
			}

			select {
			case <-ticker.C:
			case <-ctx.Done():
			}
		}
	}
}
