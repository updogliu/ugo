package ustr

import (
	"regexp"
	"strconv"
)

func MustParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func MustParseInt64(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return x
}

func MustParseBool(s string) bool {
	x, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return x
}

// Panics on error
func Match(pattern string, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	return matched
}
