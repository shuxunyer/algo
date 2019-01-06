package _15_tree

type TreeNode struct {
	data interface{}
	left *TreeNode
	right *TreeNode
}

func NewNode(val interface{}) *TreeNode{
	return &TreeNode{data :val}
}


