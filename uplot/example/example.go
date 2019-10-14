package main

import "github.com/updogliu/ugo/uplot"

func main() {
	// xs := []float64{1, 2, 8, 9, 15}
	// ys := []float64{5, 1, 2, 10, 20}
	// chart := uplot.Scatter("My Scatter", xs, ys)

	// xs := []float64{1, 2, 8, 9, 15}
	// ys := []float64{5, 1, 2, 10, 20}
	// chart := uplot.Line("My Line", xs, ys)

	chart := uplot.Line("My Dots", []struct{ X, Y float64 }{
		{1, 5}, {2, 1}, {8, 2}, {9, 10}, {15, 20},
	})

	uplot.RenderInBrowser_DestroyChart(chart)
}
