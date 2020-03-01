package utime

import "time"

// A ticker gives a instant tick in the beginnig.
type ReadyTicker struct {
	C <-chan time.Time

	stopC chan struct{}
}

func NewReadyTicker(d time.Duration) *ReadyTicker {
	c := make(chan time.Time, 1)
	c <- time.Now()

	rt := &ReadyTicker{
		C:     c,
		stopC: make(chan struct{}, 1),
	}

	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case tick := <-ticker.C:
				c <- tick
			case <-rt.stopC:
				return
			}
		}
	}()

	return rt
}

// Stop turns off a ticker. After Stop, no more ticks will be sent.
//
// Similar to `time.Ticker.Stop()`, `Stop` does not close the channel, to prevent a concurrent
// goroutine reading from the channel from seeing an erroneous "tick".
func (rt *ReadyTicker) Stop() {
	rt.stopC <- struct{}{}
}
