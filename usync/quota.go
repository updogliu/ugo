package usync

import "sync"

type Quota struct {
	*sync.Mutex
	cur int64
}

// Negative `initQty` means in debt at the beginning.
func NewQuota(initQty int64) *Quota {
	return &Quota{
		Mutex: new(sync.Mutex),
		cur:   initQty,
	}
}

// Try to acquire `qty`. When the remaining is less than `qty`, acquire all the remaining if
// `partial` is true, or acquire nothing if `partial` is false.
//
// Returns the quantity acquired. Never blocks.
//
// Precondition: `qty >= 0`
func (q *Quota) TryAcquire(qty int64, allowPartial bool) int64 {
	if qty < 0 {
		panic("Acquiring with negative qty")
	}

	q.Lock()
	defer q.Unlock()

	if q.cur <= 0 {
		return 0
	}

	var acquired int64
	if q.cur >= qty {
		acquired = qty
	} else if allowPartial {
		acquired = q.cur
	} else {
		acquired = 0
	}

	q.cur -= acquired
	return acquired
}

func (q *Quota) Release(qty int64) {
	q.Lock()
	q.cur += qty
	q.Unlock()
}
