package main

func main() {
	plotX := []float64{}
	plotY := []float64{}
	a := []int{1, 2, 3, 4}
	for i := 0; i < 100; i++ {
		op = 0
		a = append(a, i)
		elements := len(a)
		mergeSort(a, elements)
		plotX = append(plotX, float64(elements))
		plotY = append(plotY, float64(op))
	}
	drawChart(plotX, plotY)
}
