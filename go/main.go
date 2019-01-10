package main

import "fmt"

func main() {
	arr :=[]int{1,3,5,7}
	fmt.Println(arr)
	sWap(arr,1,2)
	fmt.Println(arr)

}
func sWap(arr []int,i,j int){
	arr[i], arr[j] = arr[j], arr[i]
}
