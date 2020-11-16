package main

import "fmt"

func main(){

	//append对重复数据和非重复数据的处理是相同的

	slices := []int{}

	slices = append(slices,1)
	slices = append(slices,2)
	slices = append(slices,3)

	fmt.Printf("len = %d, cap = %d\n", len(slices), cap(slices))
	fmt.Println("slices = ", slices)

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println("s2 = ", s2)
}