package uplot

import (
	"bytes"

	ec "github.com/go-echarts/go-echarts/charts"
	"github.com/pkg/browser"
	"github.com/updogliu/ugo/ulog"
)

// `args` can have one of the following types
//   - [][2]float64
//   - []struct{X, Y float64}
//   - ([]float64, []float64)
func Scatter(title string, args ...interface{}) *ec.RectChart {
	xys := getXYs(args...)

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

// `args` can have one of the following types
//   - [][2]float64
//   - []struct{X, Y float64}
//   - ([]float64, []float64)
func Line(title string, args ...interface{}) *ec.RectChart {
	xys := getXYs(args...)

	chart := ec.NewLine()
	chart.SetGlobalOptions(
		ec.TitleOpts{
			Title: title,
			TitleStyle: ec.TextStyleOpts{FontSize: 1},
		},
		ec.XAxisOpts{Type: "value", Show: true},
		ec.YAxisOpts{Type: "value", Show: true},
		ec.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}},
		ec.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}},
		ec.ToolboxOpts{Show: true},
	)
	chart.AddYAxis(title, xys)

	return &chart.RectChart
}

func getXYs(args ...interface{}) (xys [][2]float64) {
	switch v := args[0].(type) {
	case [][2]float64:
		return v

	case []struct{ X, Y float64 }:
		for _, point := range v {
			xys = append(xys, [2]float64{point.X, point.Y})
		}
		return

	case []float64:
		if len(args) != 2 {
			ulog.Panic("Unexpected num of args: ", len(args))
		}
		ys := args[1].([]float64)
		if len(v) != len(ys) {
			ulog.Panicf("xs and ys have different lengths: %v vs %v", len(v), len(ys))
		}
		for i := range v {
			xys = append(xys, [2]float64{v[i], ys[i]})
		}
		return

	default:
		ulog.Panicf("Unexpected type of the first arg: %T", args[0])
	}
	panic("unreachable")
}

// `chart` will become unusable after this function.
func RenderInBrowser_DestroyChart(chart *ec.RectChart) {
	buf := new(bytes.Buffer)
	chart.Render(buf)
	browser.OpenReader(buf)
}
