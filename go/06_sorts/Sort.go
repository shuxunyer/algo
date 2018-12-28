package _6_sorts

/*
冒泡排序、插入排序、选择排序
 */

//冒泡排序
func BubbleSort(arr []int, n int) {
	if n < 1 {
		return
	}
	for i := 0; i < n; i++ {
		flag := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

//插入排序
func InsertionSort(arr []int, n int) {
	if n < 1 {
		return
	}
	for i := 1; i < n; i++ {
		val := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = val
	}
}
