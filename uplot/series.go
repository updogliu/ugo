package uplot

import (
	"strings"

	ec "github.com/go-echarts/go-echarts/charts"
)

type Series struct {
	Name string
	Dots [][2]float64
}

type SeriesBuilder struct {
	seriesList []*Series
}

func NewSeriesBuilder() *SeriesBuilder {
	return &SeriesBuilder{}
}

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
