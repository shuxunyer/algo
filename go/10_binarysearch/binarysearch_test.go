package _10_binarysearch

import "testing"

func TestBinarySearchFirst(t *testing.T){
	a := []int{1, 3,3,5, 5,5,5, 6, 8}
	t.Log(BinarySearchFirst(a,1))
	t.Log(BinarySearchFirst(a,3))
	t.Log(BinarySearchFirst(a,4))
	t.Log(BinarySearchFirst(a,5))
	t.Log(BinarySearchFirst(a,8))
}

func TestBinarySearchLast(t *testing.T){
	a := []int{1, 3,3,5, 5,5,5, 6, 7,7,7,8}
	t.Log(BinarySearchLast(a,1))
	t.Log(BinarySearchLast(a,3))
	t.Log(BinarySearchLast(a,4))
	t.Log(BinarySearchLast(a,5))
	t.Log(BinarySearchLast(a,7))
	t.Log(BinarySearchLast(a,8))
	t.Log(BinarySearchLast(a,9))
}

func TestBinarySearchFirstGT(t *testing.T){
	a := []int{1, 3,3,5, 5,5,5, 6, 7,7,7,8}
	t.Log(BinarySearchFirstGT(a,0))
	t.Log(BinarySearchFirstGT(a,1))
	t.Log(BinarySearchFirstGT(a,3))
	t.Log(BinarySearchFirstGT(a,4))
	t.Log(BinarySearchFirstGT(a,5))
	t.Log(BinarySearchFirstGT(a,7))
	t.Log(BinarySearchFirstGT(a,8))
	t.Log(BinarySearchFirstGT(a,9))
}

func TestBinarySearchLastGT(t *testing.T){
	a := []int{1, 3,3,5, 5,5,5, 6, 7,7,7,8}
	t.Log(BinarySearchLastGT(a,0))
	t.Log(BinarySearchLastGT(a,1))
	t.Log(BinarySearchLastGT(a,3))
	t.Log(BinarySearchLastGT(a,4))
	t.Log(BinarySearchLastGT(a,5))
	t.Log(BinarySearchLastGT(a,7))
	t.Log(BinarySearchLastGT(a,8))
	t.Log(BinarySearchLastGT(a,9))
}


