package ustr

import (
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestMatch(t *testing.T) {
	r.True(t, Match("hi fox jump", "fo.*mp"))
	r.False(t, Match("hi fox jump", "Fo.*mp"))
	r.True(t, IMatch("hi fox jump", "Fo.*mp"))
	r.True(t, IMatch("hi fox jump", "Fo.*Mp"))
	r.False(t, IMatch("hi fox jump", "^Fo.*Mp"))
	r.True(t, IMatch("fox jump", "^Fo.*Mp"))
}
