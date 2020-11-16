package main

import "fmt"

func main(){
	srcSlice := []int{1,2}
	dstSlice := []int{6,6,6,6,6}
	i := copy(dstSlice, srcSlice)
	fmt.Println(i)
	fmt.Println("dst = ", dstSlice)
	fmt.Println("src = ", srcSlice)
}