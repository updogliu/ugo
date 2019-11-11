package main

import (
	"io/ioutil"

	"github.com/pkg/browser"
	"github.com/updogliu/ugo/uplot"
)

func main() {
	browser.Stdout = ioutil.Discard

	SeriesBuilder()
}

func SimpleChart() {
	chart := uplot.Scatter(
		//chart := uplot.Line(
		"My Dots", [][2]float64{{1, 5}, {2, 1}, {8, 2}, {9, 10}, {15, 20}},
		"Your Dots", []struct{ X, Y float64 }{{1, 1}, {2, 8}, {5, 10}, {9, 3}, {12, 5}},
		"Our Dots", [2][]float64{{2, 4, 8, 11, 13}, {3, 5, 2, 15, 12}},
	)
	uplot.SetSymbolSize(chart, 5)
	uplot.SetYAxisLimits(chart, -15, 30)
	uplot.RenderInBrowser_DestroyChart(chart)
}

func SeriesBuilder() {
	builder := uplot.NewSeriesBuilder()

	builder.AddDot("My Dots", 1, 5)
	builder.AddDot("Your Dots", 1, 1)
	builder.AddDot("My Dots", 2, 1)
	builder.AddDot("Your Dots", 2, 8)
	builder.AddDot("My Dots", 8, 2)
	builder.AddDot("Your Dots", 5, 3)

	chart := builder.Build("Line")
	uplot.RenderInBrowser_DestroyChart(chart)
}
