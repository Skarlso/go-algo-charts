package main

var op int

func quicksort(a []int, begin, end int) []int {
	op++
	if begin < end {
		p := partition(a, begin, end)
		quicksort(a, begin, p-1)
		quicksort(a, p+1, end)
	}
	return a
}

func partition(a []int, begin, end int) int {
	pivot := a[end]
	i := begin
	for j := begin; j < end; j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[end], a[i] = a[i], a[end]
	return i
}
