package umath

import (
	"math"
	"sort"

	"gonum.org/v1/gonum/stat"
)

// For a series of `n` values, the max and the min `n * outlierFilterRate` values are filtered
// out from the mean and std-dev calculation. Returns (NaN, NaN) if all numbers are filtered out.
//
// Precondition: `outlierFilterRate` is in [0, 1.0].
func MeanStdDev(original []float64, outlierFilterRate float64) (mean, stddev float64) {
	if outlierFilterRate < 0 || outlierFilterRate > 1.0 {
		panic("Invalid outlierFilterRate")
	}

	x := make([]float64, len(original))
	copy(x, original)
	sort.Float64s(x)

	filterNum := int(float64(len(x)) * outlierFilterRate)
	if filterNum*2 >= len(x) {
		return math.NaN(), math.NaN()
	}

	x = x[filterNum : len(x)-filterNum]
	return stat.MeanStdDev(x, nil)
}

func StdScore(x, mean, stddev float64) float64 {
	return (x - mean) / stddev
}

func Mean(x []float64) float64 {
	return stat.Mean(x, nil)
}
