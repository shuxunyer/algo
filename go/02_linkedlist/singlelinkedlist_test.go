package _2_linkedlist

import (
	"testing"
)

var l *LinkedList

func init() {
	l = NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestInsertBefore(t *testing.T) {
	l.InsertBefore(l.head.next.next.next, 333)
	l.Print()
}
func TestInsertToHead(t *testing.T) {
	for i := 10; i < 20; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	t.Log(l.FindByIndex(4))
}

func TestDeleteNode(t *testing.T) {
	l.DeleteNode(l.head.next.next.next)
	l.Print()
}
