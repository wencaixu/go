package main

import "fmt"

func modify1(p *[5]int){
	(*p)[0] = 666
	fmt.Println("modify *a =", *p)
}

func main(){
	a := [5]int{1,2,34,5} //初始化
	modify1(&a) //地址传入
	fmt.Println("main:a=",a)
}