package main

import (
	"log"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	got := mergeSort(a, len(a))
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v\n", got, want)
	}
	log.Println("operations: ", op)
}

func BenchmarkMergeSort(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := len(a)
	for i := 0; i < b.N; i++ {
		mergeSort(a, l)
	}

}
