package main

import (
	"math/rand"
	"fmt"
)

func Generate_Randnum() int{
	//rand.Seed(time.Now().Unix())
	rnd := rand.Intn(400)

	fmt.Printf("rand is %v\n", rnd)

	return rnd
}

func main(){
/*	var arr []int
	for i:=0;i<10;i++{
		ele :=Generate_Randnum()
		arr =append(arr,ele)
	}
	fmt.Println(arr)*/

	fmt.Println(1/10)
}

