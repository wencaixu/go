package main

import "fmt"

func main(){

	fmt.Println("456")

	//defer延迟调用，main函数结束前调用
	defer fmt.Print("123")

	defer fmt.Println("4567")

	fmt.Println("1234")

	//如果包含多个defer，其输出按照defer的在代码中的顺序的逆顺序

}