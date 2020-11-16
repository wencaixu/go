package main

import "fmt"

func main(){
	//定义一个变量，类型为map[int]string
	var m0 map[int]string
	fmt.Println("m0",m0)

	//对于map只有len，没有cap
	fmt.Println("len = ",len(m0))

	//根据make创建
	m1 := make(map[int]string)
	fmt.Println("m1 = ",m1)
	fmt.Println("len = ",len(m1))

	//可以通过make创建，可以指定长度，知识指定了容量，里面没有数据
	m3 := make(map[int]string,2)
	m3[1] = "mike"
	m3[2] = "go"
	m3[3] = "c++"

	fmt.Println("m3 = ",m3)
	fmt.Println("len = ",len(m3))

	//初始化
	//键值是唯一的
	m4 := map[int]string{1:"mike",2:"go",3:"c++"}
	fmt.Println("m4 = ",m4)

	//map赋值
	m5 := map[int]string{1:"mike",2:"vivo"}
	//赋值
	fmt.Println("m1 = ",m5)
	m5[1] = "c++"
	m5[3] ="go"
	fmt.Println("m1 = ",m5)
}