package _8_heap

import (
	"testing"
	"fmt"
)


func TestHeapInsert(t *testing.T) {
	heap :=NewHeap(10)
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(8)
	heap.Insert(7)
	heap.Insert(9)
	heap.Insert(11)
	heap.Insert(23)
	heap.Print()
}


func TestHeapDelete(t *testing.T) {
	heap :=NewHeap(10)
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(8)
	heap.Insert(7)
	heap.Insert(9)
	heap.Insert(11)
	heap.Insert(23)
	heap.Delete(1)
	heap.Print()
}

func TestBuildHeap(t *testing.T) {
	 arr:= []int{0,4,3,8,7,9,11,23}
	arr =buildHeap(arr)
	fmt.Println(arr)
}

func TestSort(t *testing.T) {
	arr:= []int{0,4,3,8,7,9,11,23,77,12,88,21,99,1}
	arr =Sort(arr)
	fmt.Println(arr)
}
