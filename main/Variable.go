package main

import (
	"fmt"
)

func main(){
	//变量声明格式：var 变量标识 类型
	var a int
	//初始化为0
	fmt.Print(a)

	//声明多个变量
	var s,m int
	var (
		v1 int
		v2 int
	)
	fmt.Println(v1 + v2)
	s = 10
	m = 11
	fmt.Print(s + m)

	//变量初始化
	var s1 int = 0
	var t int = 20
	fmt.Print(s1 + t)

	//自动推导类型，此时必须要初始化，通过初始化确定类型(常用)
	//出现在:=左侧变量不应该是已经声明过的变量
	c:=100
	fmt.Printf("%T",c)

	//编译器自动推导类型
	var v3 = 10
	fmt.Println(v3)

	//多重赋值
	var v4,v5,v6 int
	v4,v5,v6 = 1,2,3
	fmt.Println("多重赋值",v4,v5,v6)

	//多重复制
	i:=10
	j:=11
	fmt.Println("多重复值",i,j)

	//特殊变量名_
	_, i, _, j = 1, 2, 3, 4
	fmt.Println(i,j)

	//强制转换换
	var mean int32
	mean = int32(i)
	fmt.Println(mean)

	//常量
	const PI float32  = 3.14
	//浮点常量，自动推导类型
	const  zero   = 0.0

	//常量组
	const (
		size int64 = 1024
		eof = -1
	)

	const u,v float32  = 0,3

	const(
		x = iota //x = 0
		y = iota //y = 1
		z = iota //z = 2
		k        //隐式k = iota k = 3
	)

	fmt.Println(x,y,z,k)

	const name  = iota //遇到const时，iota 被重置

	fmt.Println(name)
}
