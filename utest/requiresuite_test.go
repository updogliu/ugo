package utest_test

import (
	"testing"

	"github.com/updogliu/ugo/utest"
)

//------------------- Example of Require and Assert ----------------------------

type MySuite struct {
	utest.RequireSuite
}

// The hook of `go test`
func TestRun_MySuite(t *testing.T) {
	utest.Run(t, new(MySuite))
}

// Uncomment and run the tests below to see the difference

// func (t *MySuite) TestUsingRequire() {
// 	t.Equal(3, 3)
// 	t.Equal(3, 4)  // terminates at here
// 	t.Equal(3, 5)  // this line will not run
// }

// func (t *MySuite) TestUsingAssert() {
// 	t.Assert().Equal(3, 3)
// 	t.Assert().Equal(3, 4)
// 	t.Assert().Equal(3, 5)
// }

//----------------- Example of building a test base struct ---------------------

type TestBase struct {
	utest.RequireSuite
	name string
}

func NewTestBase(name string) *TestBase {
	return &TestBase{name: name}
}

func (tb *TestBase) GetName() string {
	tb.True(len(tb.name) > 3, "name too short")
	return tb.name
}

// `ComponentTestSuite` is very convenient to use. It has embedded methods from both `TestBase`
// and `RequireSuite`.
type ComponentTestSuite struct {
	*TestBase
}

// The hook of `go test`
func TestRun_ComponentTestSuite(t *testing.T) {
	utest.Run(t, &ComponentTestSuite{
		TestBase: NewTestBase("abcd"),
	})
}

func (t *ComponentTestSuite) TestNaming() {
	actualName := t.GetName()
	t.Equal("abcd", actualName, "wrong name")
}
