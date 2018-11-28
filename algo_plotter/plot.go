package algoplotter

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// PlotChart will create a chart from xys coordinates.
func PlotChart(xys plotter.XYs, filename, title string) error {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "Element"
	p.Y.Label.Text = "Op"

	err = plotutil.AddLinePoints(p, xys)
	if err != nil {
		return err
	}

	return p.Save(4*vg.Inch, 4*vg.Inch, filename+".png")
}
