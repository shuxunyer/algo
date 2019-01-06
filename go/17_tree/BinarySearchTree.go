package _17_tree

import "fmt"

type BST struct {
	*BinaryTree
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootVal interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	return &BST{
		NewBinaryTree(rootVal),
		compareFunc,
	}
}

func (bst *BST) Insert(val int) bool {
	root := bst.root
	if root == nil {
		fmt.Println("the binaryTree root is nil,fuck")
		return false
	}
	return bst.insert(root, val)
}

func (bst *BST) insert(treeNode *TreeNode, val interface{}) bool {
	result := bst.compareFunc(val, treeNode.data)
	//如果要插入的数据比节点的数据大，
	if result == 1 {
		//节点的右子树为空，将新数据直接插入至节点的右子树
		if treeNode.right == nil {
			treeNode.right = NewNode(val)
			return true
		} else {
			bst.insert(treeNode.right, val)
		}
	} else if result == -1 { //如果要插入的数据比节点的数据小，
		//节点的左子树为空，将新数据直接插入至节点的左子树
		if treeNode.left == nil {
			treeNode.left = NewNode(val)
			return true
		} else {
			bst.insert(treeNode.left, val)
		}
	} else {
		//fmt.Println("fuck ")
		return false
	}
	return true
}

func (bst *BST) Find(val int) *TreeNode {
	root := bst.root
	return bst.find(root, val)
}

func (bst *BST) find(treeNode *TreeNode, val interface{}) *TreeNode {
	if treeNode == nil {
		return nil
	}
	result := bst.compareFunc(val, treeNode.data)
	if result == 0 {
		return treeNode
	} else if result == 1 {
		return bst.find(treeNode.right, val)
	} else {
		return bst.find(treeNode.left, val)
	}
}

func (bst *BST) FindNodeAndParent(val int) (node, parent *TreeNode) {
	root := bst.root
	node, parent = bst.findNodeAndParent(root, nil, val)
	return
}

func (bst *BST) findNodeAndParent(treeNode *TreeNode, parentNode *TreeNode, val interface{}) (*TreeNode, *TreeNode) {
	if treeNode == nil {
		return nil, nil
	}
	result := bst.compareFunc(val, treeNode.data)
	if result == 0 {
		return treeNode, parentNode
	} else if result == 1 {
		return bst.findNodeAndParent(treeNode.right, treeNode, val)
	} else {
		return bst.findNodeAndParent(treeNode.left, treeNode, val)
	}
}

func (bst *BST) Delete(val int) bool {
	node, parent := bst.FindNodeAndParent(val)
	if node == nil {
		return false
	}
	/*if parent != nil {
		fmt.Printf("node.val:%+v,parent.val:%+v\n", node.data, parent.data)
	} else {
		fmt.Printf("node.val:%+v\n", node.data)
	}*/
	var relationNode *TreeNode
	if node.left == nil && node.right == nil { //要删除的节点没有子节点，直接将父节点中指向要删除的子节点的指针置位nil
		relationNode = nil
	} else if node.left == nil || node.right == nil { //要删除的节点只有一个子节点，更新父节点中指向要删除的子节点，让它指向要删除节点的子节点即可
		if node.left != nil {
			relationNode = node.left
		} else {
			relationNode = node.right
		}
	} else {
		/**
		要删除的节点有两个子节点
		1 我们需要找到这个节点的右子树中的最小节点，把它替换到要删除的节点上，
		2 然后删除掉这个最小节点
		 */
		leastNode := node.FindMinInRight()
		fmt.Printf("leastNode=%+v\n",leastNode.data)
		temp := leastNode
		bst.Delete(leastNode.data.(int))

		temp.left = node.left
		temp.right = node.right
		relationNode = temp
	}
	if parent !=nil{
		//判断已知的子节点是父节点的左子树还是右子树
		if parent.left == node {
			parent.left = relationNode
		} else {
			parent.right = relationNode
		}
		fmt.Println(parent.String())
	}else{
		bst.root =nil
	}
	return true

}


func (node *TreeNode) FindMinInRight() *TreeNode {
	if node == nil || node.right ==nil{
		return nil
	}
	node =node.right
	return node.findMinInRight()
}

func (node *TreeNode) findMinInRight() *TreeNode {
	if node.left == nil {
		return node
	}
	node =node.left
	return node.findMinInRight()
}



func (bst *BST) Min() *TreeNode {
	node := bst.root
	for node != nil {
		if node.left == nil {
			break
		}
		node = node.left
	}
	return node
}

func (bst *BST) Max() *TreeNode {
	node := bst.root
	for node != nil {
		if node.right == nil {
			break
		}
		node = node.right
	}
	return node
}
