package _15_tree

import "fmt"

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(val interface{}) *BinaryTree {
	return &BinaryTree{NewNode(val)}
}

func (binaryTree *BinaryTree) InOrderTraverse() {
	p := binaryTree.root
	s := NewArrayStack()
	for p != nil || !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.left
		} else {
			temp := s.Pop().(*TreeNode)
			fmt.Printf("%+v", temp.data)
			p = temp.right
		}
	}
	fmt.Println()
}

func (binaryTree *BinaryTree) PreOrderTraverse() {
	p := binaryTree.root
	s := NewArrayStack()
	for p != nil || !s.IsEmpty() {
		if p!= nil {
			fmt.Printf("%+v", p.data)
			s.Push(p)
			p = p.left
		} else {
			p = s.Pop().(*TreeNode).right
		}
	}
	fmt.Println()
}

//error
func (binaryTree *BinaryTree) PostOrderTraverse() {
	p := binaryTree.root
	s := NewArrayStack()
	for p != nil || !s.IsEmpty() {
		if p!= nil {
			s.Push(p)
			p = p.left
		} else {
			fmt.Printf("%+v", p.data)
			p = s.Pop().(*TreeNode).right
		}
	}
	fmt.Println()
}
