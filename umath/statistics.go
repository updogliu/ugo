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
	if outlierFilterRate == 0 {
		return stat.MeanStdDev(original, nil)
	}

	x := make([]float64, len(original))
	copy(x, original)
	sort.Float64s(x)

	cutNum := int(float64(len(x)) * outlierFilterRate)
	if cutNum*2 >= len(x) {
		return math.NaN(), math.NaN()
	}

	x = x[cutNum : len(x)-cutNum]
	return stat.MeanStdDev(x, nil)
}

func GetExtremes(x []float64) (min, max float64) {
	min, max = math.Inf(+1), math.Inf(-1)
	for _, v := range x {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func Mean(x []float64) float64 {
	return stat.Mean(x, nil)
}

func StdScore(x, mean, stddev float64) float64 {
	return (x - mean) / stddev
}

func StdPercentile(x, mean, stddev float64) float64 {
	stdScore := StdScore(x, mean, stddev)
	return math.Erf(math.Abs(stdScore) / math.Sqrt2)
}
