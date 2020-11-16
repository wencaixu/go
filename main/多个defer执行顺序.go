package main

import "fmt"

func testMultiDefer(s int) int{
	return s
}

//按照书写顺序的反顺序执行
func main(){

	defer fmt.Println("123")

	defer testMultiDefer(1)

	defer fmt.Println("456")
}