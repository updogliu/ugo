package ustr

import (
	"regexp"
	"strings"
)

func SameLower(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

func Has(a, sub string) bool {
	return strings.Contains(a, sub)
}

// Case insensitive version of `Has`.
func IHas(a, sub string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(sub))
}

// Panics on invalid regular-expression.
func Match(str, reg string) bool {
	matched, err := regexp.MatchString(reg, str)
	if err != nil {
		panic(err)
	}
	return matched
}

// Case-insensitive version of `Match`.
func IMatch(str, reg string) bool {
	if !strings.HasPrefix(reg, "(?i)") {
		reg = "(?i)" + reg
	}
	return Match(str, reg)
}
