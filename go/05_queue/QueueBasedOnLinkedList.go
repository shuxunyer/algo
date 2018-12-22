package _5_queue

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedList struct {
	head   *ListNode
	tail   *ListNode
	length uint
}

func NewLinkedListQueue() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (this *LinkedList) EnQueue(val interface{}) bool {
	node := &ListNode{val, nil}
	if this.tail == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
	return true
}

func (this *LinkedList) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	val := this.head.val
	this.head = this.head.next
	this.length --
	return val
}


func (this *LinkedList) String() string {
	if this.head == nil {
		return "empty queue"
	}
	result := "head"
	for cur := this.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
