package uplot

import (
	"bytes"
	"io"

	ec "github.com/go-echarts/go-echarts/charts"
	"github.com/pkg/browser"
	"github.com/updogliu/ugo/ulog"
)

// `args` has format []{title1, xys1, title2, xys2, ...}.
// Each `xys` can have one of the following types:
//   - [][2]float64
//   - []struct{X, Y float64}
//   - [2][]float64
func Scatter(args ...interface{}) *ec.RectChart {
	chart := ec.NewScatter()
	chart.SetGlobalOptions(
		ec.XAxisOpts{Type: "value", Show: true},
		ec.YAxisOpts{Type: "value", Show: true},
		ec.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}},
		ec.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}},
		ec.ToolboxOpts{Show: true},
	)

	if len(args)%2 != 0 {
		ulog.Panic("Invalid len(args): ", len(args))
	}
	for i := 0; i < len(args); i += 2 {
		title := args[i].(string)
		xys := getXYs(args[i+1])
		chart.AddYAxis(title, xys)
	}
	return &chart.RectChart
}

// `args` has format []{title1, xys1, title2, xys2, ...}.
// Each `xys` can have one of the following types:
//   - [][2]float64
//   - []struct{X, Y float64}
//   - [2][]float64
func Line(args ...interface{}) *ec.RectChart {
	chart := ec.NewLine()
	chart.SetGlobalOptions(
		ec.XAxisOpts{Type: "value", Show: true},
		ec.YAxisOpts{Type: "value", Show: true},
		ec.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}},
		ec.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}},
		ec.ToolboxOpts{Show: true},
	)

	if len(args)%2 != 0 {
		ulog.Panic("Invalid len(args): ", len(args))
	}
	for i := 0; i < len(args); i += 2 {
		title := args[i].(string)
		xys := getXYs(args[i+1])
		chart.AddYAxis(title, xys)
	}

	return &chart.RectChart
}

func SetYAxisLimits(chart *ec.RectChart, min, max float64) {
	var yAxisOpt ec.YAxisOpts
	if len(chart.YAxisOptsList) > 0 {
		yAxisOpt = chart.YAxisOptsList[0]
	}
	yAxisOpt.Min = min
	yAxisOpt.Max = max
	chart.SetGlobalOptions(yAxisOpt)
}

func SetSymbolSize(chart *ec.RectChart, size float32) {
	for i := range chart.Series {
		chart.Series[i].SymbolSize = size
	}
}

func getXYs(arg interface{}) (xys [][2]float64) {
	switch v := arg.(type) {
	case [][2]float64:
		return v

	case []struct{ X, Y float64 }:
		for _, point := range v {
			xys = append(xys, [2]float64{point.X, point.Y})
		}
		return

	case [2][]float64:
		if len(v[0]) != len(v[1]) {
			ulog.Panicf("xs and ys have different lengths: %v vs %v", len(v[0]), len(v[1]))
		}
		for i := range v[0] {
			xys = append(xys, [2]float64{v[0][i], v[1][i]})
		}
		return

	default:
		ulog.Panicf("Unexpected type of xys arg: %T", arg)
	}
	panic("unreachable")
}

type Render interface {
	Render(w ...io.Writer) error
}

// `chart` will become unusable after this function.
func RenderInBrowser_DestroyChart(chart Render) {
	buf := new(bytes.Buffer)
	chart.Render(buf)
	browser.OpenReader(buf)
}
