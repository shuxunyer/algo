package _4_stack

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (stack *ArrayStack) Push(v interface{}) {
	if stack.top < 0 {
		stack.top = 0
	} else {
		stack.top +=1
	}
	if stack.top > len(stack.data)-1 {
		stack.data = append(stack.data, v)
	} else {
		stack.data[stack.top] = v
	}
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	val := stack.data[stack.top]
	stack.top -=1
	return val
}

func (stack *ArrayStack) IsEmpty() bool {
	if stack.top < 0 {
		return true
	}
	return false
}

func (stack *ArrayStack) Flush() {
	stack.top = -1
}

func (stack *ArrayStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

func (stack *ArrayStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		top := stack.top
		fmt.Print(stack.data[top])
		top -=1
		for top < 0 {
			fmt.Print("-->", stack.data[top])
			top -=1
		}
		fmt.Println()
	}
}
