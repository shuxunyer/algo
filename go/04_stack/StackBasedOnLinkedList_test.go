package _4_stack

import (
	"testing"
)
var list  *LinkedListStack
func init(){
	list =NewLinkedListStack()
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Print()
}

func TestLinkedListStack_Pop(t *testing.T) {
	t.Log(list.Pop())
	t.Log(list.Pop())
	t.Log(list.Pop())
	t.Log(list.Pop())
}

func TestLinkedListStack_Top(t *testing.T) {
	t.Log(list.Top())
	t.Log(list.Pop())
	t.Log(list.Top())
	list.Print()
}

func TestLinkedListStack_Flush(t *testing.T) {
	list.Flush()
	list.Print()
}


