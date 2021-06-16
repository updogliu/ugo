package usync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuota(t *testing.T) {
	r := require.New(t)

	q := NewQuota(10)
	r.Equal(int64(3), q.TryAcquire(3, false))
	r.Equal(int64(0), q.TryAcquire(8, false))
	r.Equal(int64(7), q.TryAcquire(8, true))
	r.Equal(int64(0), q.TryAcquire(2, true))
	q.Release(5)
	r.Equal(int64(2), q.TryAcquire(2, true))

	q = NewQuota(-3)
	r.Equal(int64(0), q.TryAcquire(1, true))
	r.Equal(int64(0), q.TryAcquire(1, false))
	r.Equal(int64(0), q.TryAcquire(0, true))
	r.Equal(int64(0), q.TryAcquire(0, false))
	q.Release(3)
	r.Equal(int64(0), q.TryAcquire(1, true))
	r.Equal(int64(0), q.TryAcquire(1, false))
	r.Equal(int64(0), q.TryAcquire(0, true))
	r.Equal(int64(0), q.TryAcquire(0, false))
	q.Release(2)
	r.Equal(int64(0), q.TryAcquire(5, false))
	r.Equal(int64(2), q.TryAcquire(5, true))
}
