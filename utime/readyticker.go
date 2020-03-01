package utime

import "time"

// A ticker gives a instant tick in the beginnig.
type ReadyTicker struct {
	C      <-chan time.Time
	ticker *time.Ticker
}

func NewReadyTicker(d time.Duration) *ReadyTicker {
	c := make(chan time.Time, 1)
	c <- time.Now()

	rt := &ReadyTicker{
		C:      c,
		ticker: time.NewTicker(d),
	}

	go func() {
		for tick := range rt.ticker.C {
			c <- tick
		}
	}()

	return rt
}
