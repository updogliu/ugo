package ulog

import (
	"testing"

	"github.com/updogliu/ugo/utime"
	"github.com/updogliu/ugo/utest"
)

type UtimeTestSuite struct {
	utest.RequireSuite
}

func TestRun_UtimeTestSuite(t *testing.T) {
	utest.Run(t, new(UtimeTestSuite))
}

func (t *UtimeTestSuite) TestDecodeTime() {
	tm, err := DecodeTime("190925-03:53:15.238 UTC")
	t.NoError(err)
	t.Equal(int64(1569383595238), utime.TimeToMs(tm))
}
