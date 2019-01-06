package _17_tree

import "fmt"

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(val interface{}) *BinaryTree {
	return &BinaryTree{NewNode(val)}
}

func (binaryTree *BinaryTree) InOrderTraverse() {
	if binaryTree.root ==nil{
		fmt.Println("binaryTree is nil")
		return
	}
	inOrder(binaryTree.root)
	fmt.Println()
}

func inOrder(node *TreeNode){
	if node ==nil{
		return
	}
	inOrder(node.left)
	fmt.Printf("%+v ", node.data)
	inOrder(node.right)
}

func (binaryTree *BinaryTree) PreOrderTraverse() {
	if binaryTree.root ==nil{
		fmt.Println("binaryTree is nil")
		return
	}
	preOrder(binaryTree.root)
	fmt.Println()
}

func preOrder(node *TreeNode){
	if node ==nil{

		return
	}
	fmt.Printf("%+v", node.data)
	preOrder(node.left)
	preOrder(node.right)
}

func (binaryTree *BinaryTree) PostOrderTraverse() {
	if binaryTree.root ==nil{
		fmt.Println("binaryTree is nil")
		return
	}
	postOrder(binaryTree.root)
	fmt.Println()
}

func postOrder(node *TreeNode){
	if node ==nil{
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Printf("%+v", node.data)
}