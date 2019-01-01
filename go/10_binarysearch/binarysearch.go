package _10_binarysearch

//查找第一个等于给定值的元素
func BinarySearchFirst(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < val {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//查找最后一个等于给定值的元素
func BinarySearchLast(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == n-1 || arr[mid+1] > val {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//查找第一个大于等于给定值的元素
func BinarySearchFirstGT(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] >= val {
			if mid ==0 || arr[mid-1]<val{
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//查找最后一个小于等于给定值的元素
func BinarySearchLastGT(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] <= val {
			if mid ==n-1 || arr[mid+1]>val{
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}