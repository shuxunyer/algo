package _17_tree

import "testing"

var compareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)

	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}
func TestBST_Insert(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	bst.InOrderTraverse()
}

func TestBST_Find(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)
	t.Log(bst.Find(3))
}



func TestBST_Min(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Min())
}

func TestBST_Max(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Max())
}

/**
删除的测试
	1 只有根节点 TestBST_Delete1
	2 被删除的节点没有子节点 TestBST_Delete2
	3 被删除的节点只有一个子节点 TestBST_Delete3(右子节点)  TestBST_Delete4(左子节点)
	4 被删除的节点有两个子节点 TestBST_Delete5
 */
func TestBST_Delete1(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.InOrderTraverse()
	t.Log(bst.Delete(10))
	bst.InOrderTraverse()
}

func TestBST_Delete2(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(9)
	bst.Insert(11)
	bst.Insert(8)
	bst.Insert(7)
	bst.InOrderTraverse()
	t.Log(bst.Delete(11))
	bst.InOrderTraverse()
}

func TestBST_Delete3(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(9)
	bst.Insert(12)
	bst.Insert(13)
	bst.Insert(8)
	bst.Insert(7)
	bst.InOrderTraverse()
	t.Log(bst.Delete(12))
	bst.InOrderTraverse()
}

func TestBST_Delete4(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(9)
	bst.Insert(12)
	bst.Insert(11)
	bst.Insert(8)
	bst.Insert(7)
	bst.InOrderTraverse()
	t.Log(bst.Delete(12))
	bst.InOrderTraverse()
}

func TestBST_Delete5(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(9)
	bst.Insert(8)
	bst.Insert(7)

	bst.Insert(20)
	bst.Insert(11)
	bst.Insert(30)
	bst.Insert(35)
	bst.Insert(27)
	bst.Insert(26)
	bst.Insert(28)
	bst.InOrderTraverse()
	t.Log(bst.Delete(20))
	bst.InOrderTraverse()
}


