package sort

import "testing"

func TestMergeSort(t *testing.T) {

	arr := []int{5,1,6,44,34,56,3,4, 4,11,23, 3,0, 2}
	MergeSort(arr)
	t.Log(arr)
}
