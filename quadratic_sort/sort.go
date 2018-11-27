package main

var op int

func quadraticSort(a []int) []int {
	for i := range a {
		for j := range a {
			op++
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
