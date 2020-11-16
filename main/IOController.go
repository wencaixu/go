package main

import(
	"fmt"
)

func main(){
	fmt.Println("IO 操作")

	fmt.Println("输入字符串:")
	var v int32
	fmt.Scanf("%d",&v)

	fmt.Scan(&v)

	fmt.Println(v)

	//Go语法输出
	var s int32 = 100
	fmt.Printf("%T",s)

	//整数
	a:=12
	fmt.Printf("a=%b\n",a) //%b 输出整数的二进制
	fmt.Printf("%%\n")     //输出%

	//字符类型
	var ch byte
	ch = 'a'
	fmt.Printf("%c",ch)    //输出字符a

	//浮点型
	f:=3.14
	fmt.Printf("ch = %f\n",f) //输出字符

	//复数类型
	//v:=complex(3.14,3.14)
	fmt.Printf("%v",v)

	//输出布尔类型
	fmt.Printf("%t %t\n",true,true)

	//字符串
	str:="hello go"
	fmt.Printf("str = %s\n",str)       //输出字符串
}