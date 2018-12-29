package _7_sorts

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

//从小到大排序
func merge(arr []int, start, mid, end int) {
	var (
		i       = start
		j       = mid + 1
		k       = 0
		arrCopy = make([]int, end-start+1)
	)
	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			arrCopy[k] = arr[i]
			i++
		} else {
			arrCopy[k] = arr[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		arrCopy[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		arrCopy[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], arrCopy)
}
