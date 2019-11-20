package ustr

import "strings"

func SameLower(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

func LowerHas(a, sub string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(sub))
}
