package main

import (
	"fmt"
)

func main(){

	var boolean = true

	fmt.Print(boolean)

	v1 := false
	fmt.Println(v1)

	/*	var b bool
	b = bool(1)*/

	//浮点类型
	var f1 float32
	f1 = 12
	fmt.Println("浮点类型",f1)

	//字符类型
	var ch1,ch2,ch3 byte
	ch1 = 'a'
	ch2 = 97
	ch3 = '\n'
	fmt.Printf("ch1 = %d, ch2 = %c, %c", ch1, ch2, ch3)

	//unicode编码
	var ch4 rune
	ch4 = '中'
	fmt.Printf("%c",ch4)

	//字符串变量
	var s string
	s = "123456"

	ch:=s[0]

	fmt.Printf("%c,str=%s,len=%d\n",ch,s,len(s))

	//Raw字符串
	str2:=`hello
    	mike \n \r测试
    `
	fmt.Println(str2)

	//复数类型
	var s1 complex64  //由两个float32构成的复数
	s1 = 3.2 + 12i

	s2:=complex(3.2,4)
	fmt.Println(s1,s2)

	fmt.Println(real(s1),imag(s1))

	var bools  = (2 >= 3)
	fmt.Println(bools)
}