package main

import "fmt"

func main() {
	a := []int{6, 4, 3, 5, 7, 4, 2, 4, 6, 8, 9, 3}
	sorted := mergeSort(a, len(a))
	fmt.Println("sorted: ", sorted)
}
