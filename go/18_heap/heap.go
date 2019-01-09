package _8_heap

import "fmt"

type Heap struct {
	data     []int
	capacity int
	count    int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		data:     make([]int, capacity+1),
		capacity: capacity,
		count:    0,
	}
}

func (this *Heap) Insert(val int) bool {
	if this.capacity == this.count {
		return false
	}
	this.count++
	this.data[this.count] = val
	child := this.count
	parent := this.count / 2
	for parent > 0 && this.data[child] > this.data[parent] {
		this.swap(child, parent)
		child = parent
		parent = child / 2
	}
	//fmt.Printf("child=%v,parent=%v\n",child,parent)
	return true
}

func (this *Heap) Delete(index int) int {
	if this.count == 0 {
		return -1
	}
	returnData := this.data[index]
	this.data[index] = this.data[this.count]
	this.count --

	parent := index
	for 2*parent <= this.count {
		maxIndex := parent
		if 2*parent <= this.count && this.data[2*parent] > this.data[parent] {
			maxIndex = 2 * parent
		}
		if 2*parent+1 <= this.count && this.data[2*parent+1] > this.data[parent] {
			maxIndex = 2*parent + 1
		}
		if maxIndex == parent {
			break
		}
		this.swap(parent, maxIndex)
		//fmt.Printf("parent=%v,maxIndex=%v\n", parent, maxIndex)
		parent = maxIndex
	}
	return returnData
}

func (this *Heap) delete(index int) {

}

func (this *Heap) Print() {
	if this.count <= 1 {
		fmt.Println("empty heap")
	}
	for i := 1; i <= this.count; i++ {
		fmt.Print(this.data[i], ",")
	}
	fmt.Println()
}

func (this *Heap) swap(i, j int) {
	this.data[i], this.data[j] = this.data[j], this.data[i]
}
