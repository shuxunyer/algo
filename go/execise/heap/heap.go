package heap

import "fmt"

/**
	实现一个大堆顶
 */
type Heap struct {
	data     []int
	capacity int
	count    int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		data:     make([]int, cap+1),
		capacity: cap,
		count:    0,
	}
}

func (this *Heap) Insert(val int) bool {
	if this.count == this.capacity {
		return false
	}
	this.count++
	this.data[this.count] = val
	HeapifyDownToUp(this.data,this.count)
	return true
}

func (this *Heap) DeleteTop() bool {
	if this.count == 0 {
		return false
	}
	swap(this.data, 1, this.count)
	this.count--
	HeapifyUpToDown(this.data,1,this.count)
	return true
}

func buildHeap(arr []int) []int {
	count := len(arr)-1
	if count <= 1 {
		return arr
	}
	for i := count/2; i >=1; i-- {
		HeapifyUpToDown(arr,i,count)
	}
	return arr
}

func  HeapifyUpToDown(arr []int,index,count int) {
	parent := index
	for 2*parent <= count {
		maxIndex := parent
		if 2*parent <= count && arr[maxIndex] < arr[2*parent] {
			maxIndex = 2 * parent
		}
		if 2*parent+1 <= count && arr[maxIndex] < arr[2*parent+1] {
			maxIndex = 2*parent + 1
		}
		if maxIndex == parent {
			break
		}
		swap(arr, parent, maxIndex)
		fmt.Println(arr)
		parent = maxIndex
	}
}

func HeapifyDownToUp(arr []int,index int) {
	child := index
	parent := index / 2
	for parent > 0 && arr[parent] < arr[child] {
		swap(arr, child, parent)
		child = parent
		parent = parent / 2
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (this *Heap) String() {
	if this.count < 1 {
		fmt.Println("empty heap")
		return
	}
	for i := 1; i <= this.count; i++ {
		fmt.Print(this.data[i], ",")
	}
	fmt.Println()
}
