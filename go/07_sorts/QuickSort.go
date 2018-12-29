package _7_sorts

import (
	"fmt"
	"time"
)

func QuickSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	quickSort(arr, 0, arrLen-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	time.Sleep(time.Second*1)
	fmt.Printf("start=%d,pivot=%d,end=%d\n",start,pivot,end)
	quickSort(arr, start, pivot)
	quickSort(arr, pivot+1, end)
}

func partition(arr []int, low, high int) int {
	pivotV := arr[high]
	i, j := low, low
	for ; j <= high; j++ {
		if arr[j] < pivotV {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
		//fmt.Printf("i=%d ,j=%d\n",i,j)
	}
	arr[high],arr[i] = arr[i],pivotV
	return i
}
