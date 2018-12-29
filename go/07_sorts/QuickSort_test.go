package _7_sorts

import "testing"

func TestQuickSort(t *testing.T) {
	/*arr := []int{5, 4}
	QuickSort(arr)
	t.Log(arr)*/

	//arr := []int{5, 4, 3,11,33,2,11,4,5,6,9,8,7, 2, 1}
	arr := []int{11, 8, 3, 9, 7, 1, 2, 5}
	QuickSort(arr)
	t.Log(arr)
}
