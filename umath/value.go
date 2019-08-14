package umath

func Sign(x float64) float64 {
	if x > 0 {
		return 1.0
	}
	if x < 0 {
		return -1.0
	}
	return 0.0
}

func I64Abs(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -x
}
