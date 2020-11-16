package main

import "fmt"

func main(){
	var a int = 3
	if a == 3{
		fmt.Println("a == 3")
	}

	//支持一个初始化表达式，初始化子句和条件直接需要分号隔离
	if b:=3;b==4{
		fmt.Println("b == 3")
	}else {
		fmt.Println("b == 4")
	}

	//if... else if ...else
	if a == 4{
		fmt.Println("a == 3")
	}else if a>3{
		fmt.Println("a > 3")
	}else{
		fmt.Println("a < 3")
	}

	//switch (1)
	var score = 90
	switch score {
		case 90:
			fmt.Println("excellent")
		case 80:
			fmt.Println("middle")
		case 50,60,70:
			fmt.Println("so-so")
		default:
			fmt.Println("low")
	}

	//for
	var i,sum int
	for i = 1;i <= 100;i++{
		sum += i
	}
	fmt.Println("sum=",sum)

	//range
	s:="abc"
	for i := range s{
		fmt.Printf("%d,%c\n",i,s[i])
	}

	//range
	for _,c:=range s{
		fmt.Printf("%c\n",c)
	}

	//range
	for i,c:=range s{
		fmt.Printf("%d,%c\n",i,c)
	}

	fmt.Printf("++++++++++++++++++++++++++++")
	var list = [] int{1,2,3,4}
	for i,n := range list{
		fmt.Printf("%d %d\n",i,n)
	}
}
