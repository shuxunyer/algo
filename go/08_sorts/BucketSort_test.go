package _8_sorts

import (
	"testing"
	"time"
	"fmt"
	"algo/shuxun/algo/go/07_sorts"
)

func TestBucketSort(t *testing.T) {
	var arr []int
	max :=10000
	for i:=0;i<=max*10;i++{
		arr =append(arr,Generate_Randnum(max))
	}
	t.Log(arr)
	t1:=time.Now()
	BucketSort(arr,100)
	fmt.Printf("BucketSort time=%v\n",time.Now().Sub(t1))
	t.Log(arr)

	var arr2 []int
	max2 :=10000
	for i:=0;i<=max*10;i++{
		arr2 =append(arr,Generate_Randnum(max2))
	}
	t2:=time.Now()
	_7_sorts.MergeSort(arr2)
	fmt.Printf("MergeSort time=%v\n",time.Now().Sub(t2))
}


