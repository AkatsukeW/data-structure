package sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{3, 2, 100, 22, 232, 44}
	QuickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}

func BenchmarkSort(t *testing.B) {
	arr := []int{10, 20, 22, 1, 11, 3, 45, 15}
	for i := 0; i <= t.N; i++ {
		QuickSort(arr, 0, len(arr)-1)
	}
}
