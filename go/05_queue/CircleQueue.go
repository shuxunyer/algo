package _5_queue

import "fmt"

type CircleQueue struct {
	data []interface{}
	capacity uint
	head uint
	tail uint
}

func NewCircleQueue(n uint) *CircleQueue{
	if n ==0{
		return nil
	}
	return &CircleQueue{make([]interface{},n),n,0,0}
}


func (this *CircleQueue) IsEmpty() bool{
	if this.head ==this.tail {
		return true
	}
	return false
}

func (this *CircleQueue) IsFull() bool{
	if this.head ==(this.tail+1)%this.capacity{
		return true
	}
	return false
}

func (this *CircleQueue) EnQueue(val interface{}) bool{
	if this.IsFull(){
		return false
	}
	this.data[this.tail]=val
	this.tail =(this.tail+1)%this.capacity
	return true
}

func (this *CircleQueue) DeQueue() interface{}{
	if this.IsEmpty(){
		return nil
	}
	val :=this.data[this.head]
	this.head =(this.head+1)%this.capacity
	return val
}

func (this *CircleQueue) String() string{
	if this.IsEmpty(){
		return "empty queue"
	}
	result :="head"
	cur :=this.head
	for {
		result +=fmt.Sprintf("-->%+v",this.data[cur])
		cur=(cur+1)%this.capacity
		if cur ==this.tail{
			break
		}
	}
	/*for cur :=this.head;cur !=this.tail;cur=(cur+1)%this.capacity{
		result +=fmt.Sprintf("-->%+v",this.data[cur])
	}*/
	result +="-->tail"
	return result
}