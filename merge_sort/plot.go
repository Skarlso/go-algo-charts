package main

import (
	"bytes"
	"io/ioutil"

	"github.com/wcharczuk/go-chart"
)

func drawChart(xValues, yValues []float64) {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Elements",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Operations",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: xValues,
				YValues: yValues,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	ioutil.WriteFile("mergeSort.png", buffer.Bytes(), 0766)
}
