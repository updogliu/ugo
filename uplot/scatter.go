package uplot

import (
	"bytes"

	ec "github.com/go-echarts/go-echarts/charts"
	"github.com/pkg/browser"
)

// Precondition `xs` and `ys` have the same length.
func Scatter(title string, xs, ys []float64) *ec.RectChart {
	if len(xs) != len(ys) {
		panic("xs and ys have different lengths")
	}

	xys := make([][2]float64, 0, len(xs))
	for i := range xs {
		xys = append(xys, [2]float64{xs[i], ys[i]})
	}

	chart := ec.NewScatter()
	chart.SetGlobalOptions(
		ec.TitleOpts{Title: title},
		ec.XAxisOpts{Type: "value", Show: true},
		ec.YAxisOpts{Type: "value", Show: true},
		ec.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}},
		ec.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}},
		ec.ToolboxOpts{Show: true},
	)
	chart.AddYAxis(title, xys)

	return &chart.RectChart
}

func Line(title string, xs, ys []float64) *ec.RectChart {
	if len(xs) != len(ys) {
		panic("xs and ys have different lengths")
	}

	xys := make([][2]float64, 0, len(xs))
	for i := range xs {
		xys = append(xys, [2]float64{xs[i], ys[i]})
	}

	chart := ec.NewLine()
	chart.SetGlobalOptions(
		ec.TitleOpts{Title: title},
		ec.XAxisOpts{Type: "value", Show: true},
		ec.YAxisOpts{Type: "value", Show: true},
		ec.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}},
		ec.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}},
		ec.ToolboxOpts{Show: true},
	)
	chart.AddYAxis(title, xys)

	return &chart.RectChart
}

// `chart` will become unusable after this function.
func RenderInBrowser_DestroyChart(chart *ec.RectChart) {
	buf := new(bytes.Buffer)
	chart.Render(buf)
	browser.OpenReader(buf)
}
