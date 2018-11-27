package main

import "gonum.org/v1/plot/plotter"

func main() {
	plotX := []float64{}
	plotY := []float64{}
	a := []int{1, 2, 3, 4}

	for i := 0; i < 1000; i++ {
		op = 0
		a = append(a, i)
		elements := len(a)
		mergeSort(a, elements)
		plotX = append(plotX, float64(elements))
		plotY = append(plotY, float64(op))
	}
	pts := make(plotter.XYs, len(plotX))
	for i := range pts {
		pts[i].X = plotX[i]
		pts[i].Y = plotY[i]
	}
	plotChart(pts)
}
