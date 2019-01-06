package _17_tree

import "testing"

var binaryTree *BinaryTree
func init(){
	//初始化的二叉树图片见二叉树基础(上)文章中的图片
	binaryTree = NewBinaryTree("A")
	binaryTree.root.left = NewNode("B")
	binaryTree.root.right = NewNode("C")

	l1 :=binaryTree.root.left //"B"
	l1.left = NewNode("D")
	l1.right = NewNode("E")
	r1 :=binaryTree.root.right //"C"
	r1.left = NewNode("F")
	r1.right = NewNode("G")

	l2 :=l1.left//"D"
	l2.left = NewNode("H")
	l2.right = NewNode("I")
	r2 :=l1.right//"E"
	r2.left = NewNode("J")
}
func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree.InOrderTraverse()
	//H D I B J E A F C G
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree.PreOrderTraverse()
	//A B D H I E J C F G
}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree.PostOrderTraverse()
	//H I D J E B F G C A
}

