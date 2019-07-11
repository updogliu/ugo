package ustr

import "strconv"

func MustParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		panic(err)
	}
	return f
}
