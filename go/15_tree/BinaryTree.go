package _15_tree

/**
	二叉树的遍历
 */
type BinaryTree struct {
	root *Node
}

func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{NewNode(data)}
}

func (tree *BinaryTree) PreOrder() []interface{}{
	node := tree.root
	if node == nil {
		return nil
	}
	var datas []interface{}
	preOrder(node,&datas)
	return datas
}

func preOrder(node *Node,datas *[]interface{}) {
	if node == nil {
		return
	}
	*datas =append(*datas,node.data)
	preOrder(node.left,datas)
	preOrder(node.right,datas)
}


func (tree *BinaryTree) InOrder() []interface{}{
	node := tree.root
	if node == nil {
		return nil
	}
	var datas []interface{}
	inOrder(node,&datas)
	return datas
}

func inOrder(node *Node,datas *[]interface{}) {
	if node == nil {
		return
	}
	inOrder(node.left,datas)
	*datas =append(*datas,node.data)
	inOrder(node.right,datas)
}

func (tree *BinaryTree) PostOrder() []interface{}{
	node := tree.root
	if node == nil {
		return nil
	}
	var datas []interface{}
	postOrder(node,&datas)
	return datas
}

func postOrder(node *Node,datas *[]interface{}) {
	if node == nil {
		return
	}
	postOrder(node.left,datas)
	postOrder(node.right,datas)
	*datas =append(*datas,node.data)
}