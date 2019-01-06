package _17_tree

import "fmt"

type TreeNode struct {
	data interface{}
	left *TreeNode
	right *TreeNode
}

func NewNode(val interface{}) *TreeNode{
	return &TreeNode{data :val}
}

func (this *TreeNode) String() string{
	str :=""
	if this ==nil{
		return str
	}
	str +=fmt.Sprintf("nodeVal:%v",this.data)
	if this.left !=nil{
		str +=fmt.Sprintf(", node.left:%+v",this.left.data)
	}
	if this.right !=nil{
		str +=fmt.Sprintf(", node.right:%+v",this.right.data)
	}
	return str
}




