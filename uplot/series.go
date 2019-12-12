package uplot

import (
	"fmt"
	"strings"

	ec "github.com/go-echarts/go-echarts/charts"
)

type Series struct {
	Name string
	Dots [][2]float64
}

func (s *Series) Empty() bool {
	return len(s.Dots) == 0
}

func (s *Series) Last() [2]float64 {
	return s.Dots[len(s.Dots) - 1]
}

// Helper of building and rendering one or more data series.
type SeriesBuilder struct {
	seriesList []*Series
}

func NewSeriesBuilder() *SeriesBuilder {
	return &SeriesBuilder{}
}

// Add a dot to a certain series (identified by `name`). A new series is created when adding
// its first dot.
func (sc *SeriesBuilder) AddDot(name string, x, y float64) {
	i := 0
	for ; i < len(sc.seriesList); i++ {
		if sc.seriesList[i].Name == name {
			break
		}
	}
	if i == len(sc.seriesList) {
		sc.seriesList = append(sc.seriesList, &Series{Name: name})
	}
	sc.seriesList[i].Dots = append(sc.seriesList[i].Dots, [2]float64{x, y})
}

func (sc *SeriesBuilder) PrintStats() {
	for _, s := range sc.seriesList {
		fmt.Printf("[%v] Count: %v\n", s.Name, len(s.Dots))
	}
}

func (sc *SeriesBuilder) Build(typ string) *ec.RectChart {
	var data []interface{}
	for _, s := range sc.seriesList {
		data = append(data, s.Name)
		data = append(data, s.Dots)
	}

	switch strings.ToLower(typ) {
	case "line":
		return Line(data...)
	case "scatter":
		return Scatter(data...)
	default:
		panic("Invalid type: " + typ)
	}
}
