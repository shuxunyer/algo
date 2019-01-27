package _8_sorts

import (
	"algo/shuxun/algo/go/07_sorts"
	"fmt"
)

/**
	桶排序
	例如：将1-100000范围内的数据
	放入100个桶中，如果数据分布均匀，每个桶中存1000条数据
 */
func BucketSort(arr []int,internal int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	min, max := arr[0], arr[0]
	//40,11,3,5,8
	for _, ele := range arr {
		if ele < min {
			min = ele
		}
		if ele > max {
			max = ele
		}
	}
	bucketNum := (max - min) / internal
	fmt.Printf("bucketNum=%v,max=%v,min=%v\n",bucketNum,max,min)
	bucketArr :=make([][]int,bucketNum+1)
	for _, ele := range arr {
		//fmt.Printf("ele=%v,ele/interval=%v\n",ele,ele/interval)
		bucketArr[ele/internal] = append(bucketArr[ele/internal], ele)//TODO
	}

	for _, bucket := range bucketArr {
		_7_sorts.MergeSort(bucket)
	}

	sortedArr := make([]int,0, arrLen)
	for _, bucket := range bucketArr {
		fmt.Printf("sorted after=%v\n",bucket)
		sortedArr = append(sortedArr, bucket...)
	}
	fmt.Printf("all=%v\n",sortedArr)
	copy( arr,sortedArr)
}
