package main

import (
	"github.com/Skarlso/go-algo-charts/algoplotter"
	"gonum.org/v1/plot/plotter"
)

func main() {
	plotX := []float64{}
	plotY := []float64{}
	a := []int{4, 3, 2, 1}

	for i := 0; i < 1000; i++ {
		op = 0
		a = append(a, i)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		elements := len(a)
		quadraticSort(a)
		plotX = append(plotX, float64(elements))
		plotY = append(plotY, float64(op))
	}
	pts := make(plotter.XYs, len(plotX))
	for i := range pts {
		pts[i].X = plotX[i]
		pts[i].Y = plotY[i]
	}
	algoplotter.PlotChart(pts, "Quadractic Sort", "quadratic_sort")
}
