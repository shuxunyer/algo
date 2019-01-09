package _8_heap

import "testing"

var heap *Heap
func init(){
	heap =NewHeap(10)
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(8)
	heap.Insert(7)
	heap.Insert(9)
	heap.Insert(11)
	heap.Insert(23)
}

func TestHeapInsert(t *testing.T) {
	heap.Print()
}

func TestHeapDelete(t *testing.T) {
	heap.Delete(1)
	heap.Print()
}
