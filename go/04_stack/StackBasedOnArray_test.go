package _4_stack

import (
	"testing"
)

func Test1(t *testing.T){
	data:=make([]interface{}, 1, 32);
	data[0]=1
}

func TestArrayStack_Push(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	t.Log(s.Pop())
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Push(4)
	t.Log(s.Pop())
	s.Print()
}
