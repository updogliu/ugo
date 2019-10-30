package main

import "github.com/updogliu/ugo/uplot"

func main() {
	//chart := uplot.Scatter(
	chart := uplot.Line(
		"My Dots", [][2]float64{{1, 5}, {2, 1}, {8, 2}, {9, 10}, {15, 20}},
		"Your Dots", []struct{ X, Y float64 }{{1, 1}, {2, 8}, {5, 10}, {9, 3}, {12, 5}},
		"Our Dots", [2][]float64{{2, 4, 8, 11, 13}, {3, 5, 2, 15, 12}},
	)

	uplot.RenderInBrowser_DestroyChart(chart)
}
