package _8_sorts

import (
	"github.com/ethereum/go-ethereum/common/math"
)

/**
	计数排序
	1 得到计数排序数组的最大值
	2 统计数组中每个元素出现的次数
	3 统计小于等于每个元素出现的次数
	4 遍历排序数组，
 */
func CountingSort(arr []int){
	arrLen :=len(arr)
	if arrLen <=1{
		return
	}
	var (
		max int =math.MinInt32
	)
	for _,ele :=range arr{
		if ele >max{
			max =ele
		}
	}
	//fmt.Printf("max=%v\n",max)
	countArr :=make([]int,max+1)
	for _,ele :=range arr{
		countArr[ele]++
	}
	//fmt.Printf("countArr=%v\n",countArr)
	for i:=1;i<=max;i++{
		countArr[i]+=countArr[i-1]
	}
	//fmt.Printf("countArr=%v\n",countArr)

	sortedArr :=make([]int,arrLen)
	for _,ele :=range arr{
		index :=countArr[ele]-1
		sortedArr[index]=ele
		countArr[ele]--
	}
	copy(arr,sortedArr)
}
