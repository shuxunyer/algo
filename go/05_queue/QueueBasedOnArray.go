package _5_queue

import (
	"github.com/ethereum/go-ethereum/log"
	"fmt"
)

type ArrayQueue struct {
	data                 []interface{}
	capacity, head, tail uint
}

func NewArrayQueue(n uint) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (this *ArrayQueue) EnQueue(val interface{}) bool {
	if this.capacity == this.tail {
		return false
	}
	this.data[this.tail] = val
	this.tail ++
	return true
}

func (this *ArrayQueue) DeQueue() interface{} {
	if this.tail == this.head {
		log.Warn("empty queue")
		return nil
	}
	val := this.data[this.head]
	this.head++
	return val
}

func (this *ArrayQueue) String() string {
	if this.head == this.tail {
		return "empty queue"
	}
	result := "head"
	for i := this.head; i <= this.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", this.data[i])
	}
	result += "<-tail"
	return result
}
