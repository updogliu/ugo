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

func I32MakeMin(currentMin *int32, candidate int32) {
	if candidate < *currentMin {
		*currentMin = candidate
	}
}

func I32MakeMax(currentMax *int32, candidate int32) {
	if candidate > *currentMax {
		*currentMax = candidate
	}
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

func IBetween(x, min, max int) bool {
	return min <= x && x <= max
}

func I64Between(x, min, max int64) bool {
	return min <= x && x <= max
}

func F64Between(x, min, max float64) bool {
	return min <= x && x <= max
}

func IChoose(condition bool, tValue, fValue int) int {
	if condition {
		return tValue
	}
	return fValue
}

func I32Choose(condition bool, tValue, fValue int32) int32 {
	if condition {
		return tValue
	}
	return fValue
}

func I64Choose(condition bool, tValue, fValue int64) int64 {
	if condition {
		return tValue
	}
	return fValue
}

func F64Choose(condition bool, tValue, fValue float64) float64 {
	if condition {
		return tValue
	}
	return fValue
}
