package utest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// `RequireSuite` implements the `github.com/stretchr/testify/suite.TestingSuite` interface.
// Comparing to struct `github.com/stretchr/testify/suite.Suite`, `RequireSuite` embeds
// `require.Assertions`, not `assert.Assertions`. Therefore, a test in `RequireSuite` will terminate
// at the first failure.
//
// Usage:
// - `utest.Run(t, requireSuite)` to run the tests in the suite.
// - `requireSuite.Equal(1, 2)` would terminate the current test, reporting the failure.
// - `requireSuite.Assert().Equal(1, 2)` returns false indicating whether the test has failed.
// - When running tests, use -testify.m=<pattern> to only run tests with names containing <pattern>.
//
// See requiresuite_test.go for examples.
//
type RequireSuite struct {
	*require.Assertions
	assert *assert.Assertions
	t      *testing.T
}

func (rs *RequireSuite) Assert() *assert.Assertions {
	if rs.assert == nil {
		rs.assert = assert.New(rs.t)
	}
	return rs.assert
}

func (rs *RequireSuite) T() *testing.T {
	return rs.t
}

// SetT sets the current *testing.T context.
func (rs *RequireSuite) SetT(t *testing.T) {
	rs.Assertions = require.New(t)
	rs.assert = assert.New(t)
	rs.t = t
}

var Run = suite.Run
