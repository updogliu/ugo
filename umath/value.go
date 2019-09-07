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

func I64Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func I64Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func IMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns the first value `v` that `v >= x` and `v` is a multiple of `unit`.
func I64RoundUp(x, unit int64) int64 {
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
func I64RoundDown(x, unit int64) int64 {
	if unit <= 0 {
		panic("non-positive unit")
	}
	v := x - x%unit
	if v > x {
		v -= unit
	}
	return v
}

func IRoundUp(x, unit int) int {
	return int(I64RoundUp(int64(x), int64(unit)))
}

func IRoundDown(x, unit int) int {
	return int(I64RoundDown(int64(x), int64(unit)))
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

func F64MakeMin(currentMin *float64, candidate float64) {
	if candidate < *currentMin {
		*currentMin = candidate
	}
}

func F64MakeMax(currentMax *float64, candidate float64) {
	if candidate > *currentMax {
		*currentMax = candidate
	}
}
