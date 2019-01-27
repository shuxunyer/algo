package _15_tree

type Node struct {
	data interface{}
	left *Node
	right *Node
}

func NewNode(val interface{}) *Node{
	return &Node{data:val}
}
