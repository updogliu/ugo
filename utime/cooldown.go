package utime

import (
	"sync"
	"time"
)

// Concurrency-safe
type Cooldown struct {
	Duration time.Duration

	sync.Mutex
	lastReadyTime time.Time
}

// Returns a Cooldown that is ready.
func NewCooldown(duration time.Duration) *Cooldown {
	return &Cooldown{Duration: duration}
}

// Returns a Cooldown which will be ready after `duration`.
func NewUnreadyCooldown(duration time.Duration) *Cooldown {
	return &Cooldown{Duration: duration, lastReadyTime: time.Now()}
}


// Returns true if and only if one of the following is true
//   1. It is the first call of `Ready()` on `cd`.
//   2. The last time `cd.Ready()` returns true was at least `cd.Duration` ago.
func (cd *Cooldown) Ready() bool {
	cd.Lock()
	defer cd.Unlock()

	now := time.Now()
	if now.Sub(cd.lastReadyTime) >= cd.Duration {
		cd.lastReadyTime = now
		return true
	}
	return false
}
