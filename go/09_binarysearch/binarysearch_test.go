package _9_binarysearch

import "testing"

func TestBinarySearch(t *testing.T){
	a := []int{1, 3, 5, 6, 8}
	t.Log(BinarySearch(a,1))
	t.Log(BinarySearch(a,3))
	t.Log(BinarySearch(a,4))
	t.Log(BinarySearch(a,8))
}

func TestBinarySearchRecursive(t *testing.T){
	a := []int{1, 3, 5, 6, 8}
	t.Log(BinarySearchRecursive(a,1))
	t.Log(BinarySearchRecursive(a,3))
	t.Log(BinarySearchRecursive(a,4))
	t.Log(BinarySearchRecursive(a,8))
}

