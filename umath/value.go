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

func IAbs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Returns the first value `v` that `v >= x` and `v` is a multiple of `unit`.
func I64RoundUp(x int64, unit int64) int64 {
	if unit <= 0 {
		panic("non-positive unit")
	}
	v := x - x%unit
	if v < x {
		v += unit
	}
	return v
}

// Returns the last value `v` that `v <= x` and `v` is a multiple of `unit`.
func I64RoundDown(x int64, unit int64) int64 {
	if unit <= 0 {
		panic("non-positive unit")
	}
	v := x - x%unit
	if v > x {
		v -= unit
	}
	return v
}

func I64MakeMin(currentMin *int64, candidate int64) {
	if candidate < *currentMin {
		*currentMin = candidate
	}
}

func I64MakeMax(currentMax *int64, candidate int64) {
	if candidate > *currentMax {
		*currentMax = candidate
	}
}

func IMakeMin(currentMin *int, candidate int) {
	if candidate < *currentMin {
		*currentMin = candidate
	}
}

func IMakeMax(currentMax *int, candidate int) {
	if candidate > *currentMax {
		*currentMax = candidate
	}
}
