package sort

import "fmt"

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	privot := partition(arr, start, end)
	if privot == end {
		quickSort(arr, start, privot-1)
		quickSort(arr, privot, end)
	}else{
		quickSort(arr, start, privot)
		quickSort(arr, privot+1, end)
	}

}

//{6,11,3,9,8,4,7}
func partition(arr []int, start, end int) int {
	var (
		val = arr[end]
		i   = start
		j   = start
	)

	for ; j < end; j++ {
		if arr[j] < val {
			if i < j {
				swap(arr, i, j)
			}
			i++
		}
	}
	swap(arr, i, j)
	fmt.Printf("start=%v,end=%v,partition=%v\n", start, end, i)
	return i
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
