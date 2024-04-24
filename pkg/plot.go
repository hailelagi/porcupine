package pkg

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func PlotExample() {
	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		lineShowLabel(),
		lineSymbols(),
		lineMarkPoint(),
		lineSplitLine(),
		lineStep(),
		lineSmooth(),
		lineArea(),
		lineSmoothArea(),
		lineOverlap(),
		lineMulti(),
		lineDemo(),
	)
	f, err := os.Create("assets/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
