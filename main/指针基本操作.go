package main

import "fmt"

func main(){
	var a int = 10

	//变脸具有两层含义：变量地址，变量内存
	fmt.Printf("a = %d\n",a)

	fmt.Printf("&a = %v\n",&a)

	//保存某个变量的地址，需要指针类型 * int 保存int地址  ** int 保存 *int地址
	//声明（定义），定义只是特殊的声明
	//定义一个变量P，类型为*int
	var p * int
	p = &a
	fmt.Printf("p = %v,&a=%v\n",p,&a)

	//*p 操作的不是p的内存，是p所指的内存（a）
	*p = 666
	fmt.Printf("*p = %v,a = %v\n",*p,a)
}