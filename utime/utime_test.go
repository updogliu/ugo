package utime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type UtimeTestSuite struct {
	suite.Suite
}

// The hook of `go test`
func TestUtimeTestSuite(t *testing.T) {
	suite.Run(t, new(UtimeTestSuite))
}

func (t *UtimeTestSuite) TestAssertRealTime() {
	goodTimes := []time.Time{
		time.Now(),
		time.Date(2017, time.December, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2027, time.December, 01, 23, 59, 59, 999999999, time.UTC),
	}
	badTimes := []time.Time{
		time.Date(2017, time.November, 30, 23, 59, 59, 999999999, time.UTC),
		time.Date(2027, time.December, 02, 0, 0, 0, 0, time.UTC),
	}

	for _, time := range goodTimes {
		t.NotPanics(func() {
			AssertRealTimeSec(time.Unix())
			AssertRealTimeMs(time.UnixNano() / 1e6)
		})
	}
	for _, time := range badTimes {
		t.Panics(func() {
			AssertRealTimeSec(time.Unix())
			AssertRealTimeMs(time.UnixNano() / 1e6)
		})
	}
}
