package ubig

import (
	"testing"

	"github.com/updogliu/ugo/utest"
)

type TestUbigSuite struct {
	utest.RequireSuite
}

// The hook of `go test`
func TestRun_TestUbigSuite(t *testing.T) {
	utest.Run(t, new(TestUbigSuite))
}

func (t *TestUbigSuite) TestClone() {
	a := U64(123)
	a2 := a
	b := Clone(a)

	a.Add(a, U64(1))

	t.Equal(U64(124), a)
	t.Equal(U64(124), a2)
	t.Equal(U64(123), b)
}

func (t *TestUbigSuite) TestQuo() {
	t.Equal(U64(1000), Quo(U64(123056), U64(123)))
	t.Equal(U64(1001), Quo(U64(123123), U64(123)))

	t.Equal(I64(-1000), Quo(I64(-123056), I64(123)))
	t.Equal(I64(-1001), Quo(I64(-123123), I64(123)))

	t.Equal(I64(1000), Quo(I64(-123056), I64(-123)))
	t.Equal(I64(1001), Quo(I64(-123123), I64(-123)))
}
