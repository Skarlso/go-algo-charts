package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func plotChart(xys plotter.XYs) {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Quadratic Sort"
	p.X.Label.Text = "Element"
	p.Y.Label.Text = "Op"

	err = plotutil.AddLinePoints(p, xys)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "quadratic_sort.png"); err != nil {
		panic(err)
	}
}
