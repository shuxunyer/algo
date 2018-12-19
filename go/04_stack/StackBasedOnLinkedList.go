package _4_stack

import "fmt"

/*
基于链表实现的栈
author:shuxun
*/

type node struct {
	next  *node
	value interface{}
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (stack *LinkedListStack) Push(val interface{}) {
	stack.topNode = &node{stack.topNode, val}
}

func (stack *LinkedListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	val := stack.topNode
	stack.topNode = stack.topNode.next
	return val
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.topNode == nil
}

func (stack *LinkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.topNode.value
}

func (stack *LinkedListStack) Flush() {
	stack.topNode = nil
}

func (stack *LinkedListStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := stack.topNode
		fmt.Print( cur.value)
		cur =cur.next
		for cur != nil {
			fmt.Print("-->", cur.value)
			cur = cur.next
		}
		fmt.Println()
	}
}
