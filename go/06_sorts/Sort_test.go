package _6_sorts

import (
	"testing"
	"fmt"
	"time"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	BubbleSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{1,5,9,6,3,7,22,45,66,55,66,77,88,1,5,9,6,3,7,45,66,55,66,77,88,87,67,56,322,45,66,55,66,87,67,56,33,12,8,5,10,3,4,11,287,67,56,33,12,8,5,10,3,4,11,22,11,11}
	fmt.Println("排序前：",arr)
	InsertionSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestTimeBetweenInsertionSortAndBubbleSort(t *testing.T){
	arr := []int{1,5,9,6,3,7,22,45,66,55,66,77,88,1,5,9,6,3,7,45,66,55,66,77,88,87,67,56,322,45,66,55,66,87,67,56,33,12,8,5,10,3,4,11,287,67,56,33,12,8,5,10,3,4,11,22,11,11}
	t1 :=time.Now()
	for i:=0;i<10000;i++{
		 InsertionSort(arr,len(arr))
	}
	sum :=time.Now().UnixNano()-t1.UnixNano()
	t.Log("BubbleSort","time",sum)

	arr = []int{1,5,9,6,3,7,22,45,66,55,66,77,88,1,5,9,6,3,7,45,66,55,66,77,88,87,67,56,322,45,66,55,66,87,67,56,33,12,8,5,10,3,4,11,287,67,56,33,12,8,5,10,3,4,11,22,11,11}
	t2 :=time.Now()
	for i:=0;i<10000;i++{
		BubbleSort(arr,len(arr))
	}
	sum2 :=time.Now().UnixNano()-t2.UnixNano()
	t.Log("InsertionSort","time",sum2)
}
