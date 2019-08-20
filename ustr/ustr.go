package ustr

import "strings"

func SameLower(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}
