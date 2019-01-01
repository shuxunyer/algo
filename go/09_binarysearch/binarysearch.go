package _9_binarysearch

func BinarySearch(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == v {
			return mid
		} else if arr[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursive(arr []int,val int) int{
	n := len(arr)
	if n == 0 {
		return -1
	}
	return bs(arr,val,0,n-1)
}

func bs(arr []int,val int,low,high int) int{
	if low >high {
		return -1
	}
	mid :=low +(high-low)/2
	if arr[mid]==val{
		return mid
	}else if arr[mid]>val{
		return bs(arr,val,low,mid-1)
	}else{
		return bs(arr,val,mid+1,high)
	}

}
