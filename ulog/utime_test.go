package ulog

import (
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestDecodeTime(t *testing.T) {
	tm, err := DecodeTime("190925-03:53:15.238 UTC")
	r.NoError(t, err)
	r.Equal(t, int64(1569383595238000000), tm.UnixNano())
}
