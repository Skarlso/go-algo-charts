package runner

import (
	"github.com/Skarlso/go-algo-charts/algoplotter"
	"gonum.org/v1/plot/plotter"
)

// Run a single sorting algorithm in a loop and create a plot of the result.
func Run(f func([]int) int, filename, title string) {
	plotX := []float64{}
	plotY := []float64{}
	a := []int{4, 3, 2, 1}
	for i := 0; i < 2000; i++ {
		a = append(a, i)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		elements := len(a)
		op := f(a)
		plotX = append(plotX, float64(elements))
		plotY = append(plotY, float64(op))
	}
	pts := make(plotter.XYs, len(plotX))
	for i := range pts {
		pts[i].X = plotX[i]
		pts[i].Y = plotY[i]
	}
	algoplotter.PlotChart(pts, filename, title)
}
