package main

import "fmt"

//函数返回值是一个匿名函数，返回一个函数类型
func test() func() int{
	var index int
	return func() int{
		index ++
		return index
	}
}



func main(){

	//返回值是一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数
	//f来调用闭包函数
	//函数不关心捕获变量和常量是否已经超出了作用域
	//只要闭包存在，这些变量还会存在

	f := test()
	f()//1

	fmt.Println(f())//2
	fmt.Println(f())//3
	fmt.Println(f())//4
	fmt.Println(f())//5

}
