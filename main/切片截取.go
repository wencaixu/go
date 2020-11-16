package main

import "fmt"

func main(){
	array := []int{0,1,23,4,5}

	s1 := array[:]  //起始下标，长度，容量

	//s2 := array[0:2:10]

	s2 := array[0:2:3]

	fmt.Println("s1=",s1)

	fmt.Println("s1=",s2)

	fmt.Printf("len = %d,cap = %d\n",len(s1),cap(s1))

	//切片操作某个元素
	date := array[1]
	fmt.Println("data=",date)

	s3 := array[2:3:4]
	fmt.Println("s2 = ", s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[:3] //从0开始，取6个元素，容量也为6

	fmt.Println(s4)

	s5 := array[3:] //从下标3开始，到结尾,不包括第三个
	fmt.Println("s5=",s5)

}