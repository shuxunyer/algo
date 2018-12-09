package _2_linkedlist

import "fmt"

/*
单链表基本操作
author:shuxun
*/

type ListNode struct {
	next *ListNode
	val  interface{}
}

type LinkedList struct {
	head *ListNode
	len  uint
}

func NewListNode(val interface{}) *ListNode {
	return &ListNode{nil, val}
}

func (node *ListNode) GetValue() interface{} {
	return node.val
}

func (node *ListNode) GetNext() *ListNode {
	return node.next
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, val interface{}) bool {
	if p == nil {
		return false
	}
	oldNode := p.next
	newNode := &ListNode{nil, val}
	p.next = newNode
	newNode.next = oldNode
	this.len++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, val interface{}) bool {
	if nil == p || p == this.head {
		return false
	}

	pre := this.head
	cur := this.head.next
	for cur != nil {
		if p == cur {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil { //表明p节点没找到
		return false
	}
	newNode := &ListNode{nil, val}
	pre.next = newNode
	newNode.next = cur
	this.len ++
	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(val interface{}) bool {
	return this.InsertAfter(this.head, val)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(val interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, val)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index > this.len {
		return nil
	}
	cur := this.head
	var i uint = 0
	for ; i <= index; i++ {
		cur = cur.next
	}
	return cur
}

//删除节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	pre := this.head
	cur := this.head.next
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}

	pre.next = cur.next
	p = nil
	this.len --
	return true
}

func (this *LinkedList) Print() {
	format := ""
	cur := this.head.next
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.val)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
