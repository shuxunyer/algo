package _7_linkedlist

import (
	"fmt"
)

/*
单链表练习题
author:shuxun
*/

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
}

func (this *LinkedList) Print() {
	if nil == this.head {
		return
	}
	format := ""
	cur := this.head.next
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

func (this *LinkedList) test2() {
	var pre *ListNode = nil
	cur := this.head.next
	cur.next = pre
	pre = &ListNode{nil, 3}
}

/*
单链表反转
时间复杂度：O(N)
*/
func (this *LinkedList) Reverse() {
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}
	var pre *ListNode = nil
	cur := this.head.next
	for nil != cur {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}
	this.head.next = pre
}

/*
判断单链表是否有环
*/
func (this *LinkedList) HasCycle() bool {
	if this.head != nil {
		fast := this.head
		slow := this.head

		for fast != nil && fast.next != nil {
			fast = fast.next.next
			slow = slow.next
			if fast == slow {
				return true
			}
		}
	}
	return false
}

/*
两个有序单链表合并
*/
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	list := &LinkedList{&ListNode{}}
	cur := list.head
	cur1, cur2 := l1.head.next, l2.head.next
	for nil != cur1 && nil != cur2 {
		if cur1.value.(int) > cur2.value.(int) {
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}

	if cur1 == nil {
		cur.next = cur2
	}
	if cur2 == nil {
		cur.next = cur1
	}
	return list
}

/*
删除倒数第N个节点
*/
func (this *LinkedList) DeleteBottomN(n int) {
	if nil == this.head || nil == this.head.next || n <= 0 {
		return
	}
	fast := this.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}
	if fast == nil {
		return
	}
	slow :=this.head
	for fast.next !=nil{
		slow =slow.next
		fast =fast.next
	}
	fmt.Printf("slow.value =%v\n",slow.value)
	slow.next =slow.next.next
}

/*
获取中间节点
*/
func (this *LinkedList) FindMiddleNode() *ListNode {
	 if nil ==this.head || nil ==this.head.next{
	 	return nil
	 }
	 if this.head.next.next ==nil{
	 	return this.head.next
	 }

	 fast ,slow := this.head,this.head
	 for fast !=nil && fast.next !=nil{
	 	slow =slow.next
	 	fast =fast.next.next
	 }
	 return slow
}
