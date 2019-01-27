package _8_sorts

import (
	"testing"
	"math/rand"
)

func TestCountingSort(t *testing.T) {
	var arr []int
	for i:=0;i<=10000;i++{
		arr =append(arr,Generate_Randnum(400))
	}
	CountingSort(arr)
	t.Log(arr)
}

func Generate_Randnum(max int) int{
	//rand.Seed(time.Now().Unix())
	rnd := rand.Intn(max)
	//fmt.Printf("rand is %v\n", rnd)
	return rnd
}

