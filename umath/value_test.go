package umath

import (
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestI64Round(t *testing.T) {
	r.Equal(t, int64(-500), I64RoundUp(-500, 100))
	r.Equal(t, int64(-100), I64RoundUp(-180, 100))
	r.Equal(t, int64(-100), I64RoundUp(-100, 100))
	r.Equal(t, int64(0), I64RoundUp(0, 100))
	r.Equal(t, int64(100), I64RoundUp(50, 100))
	r.Equal(t, int64(1100), I64RoundUp(1001, 100))

	r.Equal(t, int64(-500), I64RoundDown(-500, 100))
	r.Equal(t, int64(-200), I64RoundDown(-180, 100))
	r.Equal(t, int64(-100), I64RoundDown(-100, 100))
	r.Equal(t, int64(0), I64RoundDown(0, 100))
	r.Equal(t, int64(0), I64RoundDown(50, 100))
	r.Equal(t, int64(1000), I64RoundDown(1001, 100))

	r.Equal(t, int64(1000), I64RoundDown(1001, 100))
}

func TestStdPercentile(t *testing.T) {
	r.Equal(t, float64(0), StdPercentile(100, 100, 20))
	r.InEpsilon(t, 0.682689492137086, StdPercentile(120, 100, 20), 1e-8)
	r.InEpsilon(t, 0.954499736103642, StdPercentile(-10, 0, 5), 1e-8)
	r.InEpsilon(t, 0.997300203936740, StdPercentile(5, -10, 5), 1e-8)
}
