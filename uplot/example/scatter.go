package main

import "github.com/updogliu/ugo/uplot"

func main() {
	xs := []float64{1, 2, 8, 9, 15}
	ys := []float64{5, 1, 2, 10, 20}

	chart := uplot.Scatter("My Scatter", xs, ys)
	uplot.RenderInBrowser_DestroyChart(chart)
}
