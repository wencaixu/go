package main

import (
	"fmt"
)

func main(){
	i := 0
	s := "mike"

	//闭包的实现方式一
	func1 := func() {
		fmt.Printf("%d %s\n",i,s)
	}

	//函数调用
	func1()

	//方式一调用方式使用函数类型
	type FuncType func()

	var f2 FuncType = func1

	f2() //函数调用

	//方式二调用
	var f3 FuncType = func(){
		fmt.Printf("%d %s\n",i,s)
	}

	f3()//函数调用

	//匿名函数，无参数无返回值
	func(){
		fmt.Printf("%d %s\n",i,s)
	}()

	//匿名函数，有参数有返回值
	v:= func(a,b int)(result int) {
		result = a + b
		return
	}(1,1)

	fmt.Println("v = ",v)

	//闭包捕获外部变量
	func(){
		i = 100
		s = "go"
		// 100 go
	    fmt.Printf("%d %s\n",i,s)
	}()
	// 100 go
	fmt.Printf("%d %s\n",i,s)

	fmt.Println("123")
}