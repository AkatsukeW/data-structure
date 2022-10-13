package sort

import (
	"testing"
)

func TestShell(t *testing.T) {
	arr := []int{3, 20, 4, 100, 76, 232, 111}
	ShellSort(arr)

	t.Log(arr)
}
