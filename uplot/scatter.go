package uplot

import (
	"io/ioutil"
	"os"
	"os/exec"

	ec "github.com/updogliu/go-echarts/charts"
	"github.com/updogliu/ugo/ulog"
	"github.com/updogliu/ugo/utime"
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

// `chart` will become unusable after this function.
func RenderInBrowser_DestroyChart(chart *ec.RectChart) {
	fout, err := ioutil.TempFile("", "uplot_scatter_*.html")
	if err != nil {
		ulog.Panic("Failed to create temp html file: ", err)
	}
	defer os.Remove(fout.Name()) // clean up

	chart.Render(fout)
	if err := exec.Command("google-chrome", fout.Name()).Run(); err != nil {
		ulog.Panicf("Failed to open %v with google-chrome: ", err)
	}
	if err := fout.Close(); err != nil {
		ulog.Panicf("Failed to close %v: %v", fout.Name(), err)
	}
	utime.SleepMs(500) // not to delete the tmp file too quickly
}
