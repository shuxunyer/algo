package _5_queue

import "testing"

func TestArrayQueue_EnQueue(t *testing.T) {
	queue :=NewArrayQueue(5)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.EnQueue(6)
	queue.EnQueue(7)
	queue.EnQueue(8)
	t.Log(queue)
}

func TestArrayQueue_DeQueue(t *testing.T) {
	queue :=NewArrayQueue(5)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.DeQueue()
	t.Log(queue)
	queue.DeQueue()
	t.Log(queue)
	queue.DeQueue()
	t.Log(queue)
	queue.DeQueue()
	t.Log(queue)
	queue.String()
}