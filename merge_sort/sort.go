package main

import (
	"math"
)

var op int

func mSort(a []int) int {
	op = 0
	mergeSort(a, len(a))
	return op
}

func merge(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	if a[0] < b[0] {
		return append([]int{a[0]}, merge(a[1:len(a)], b)...)
	}
	return append([]int{b[0]}, merge(a, b[1:len(b)])...)
}

func mergeSort(a []int, n int) []int {
	op++
	if n == 1 {
		return a
	}
	middle := int(math.Floor(float64(n / 2)))
	left := a[0:middle]
	right := a[middle:n]
	return merge(mergeSort(left, middle), mergeSort(right, n-middle))
}
