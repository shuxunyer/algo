package sort

import "testing"

func TestQuickSort(t *testing.T) {
	/*arr := []int{5, 4}
	QuickSort(arr)
	t.Log(arr)*/

	//arr := []int{5, 4, 3,11,33,2,11,4,5,6,9,8,7, 2, 1}
	arr := []int{6,11,3,9,8,4,7}
	QuickSort(arr)
	t.Log(arr)
}
