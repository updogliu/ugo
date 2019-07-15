package ustr

import (
	"regexp"
	"strconv"
)

func MustParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		panic(err)
	}
	return f
}

// Panics on error
func Match(pattern string, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	return matched
}
