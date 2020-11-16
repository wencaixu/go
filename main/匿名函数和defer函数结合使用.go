package main

import "fmt"

func main(){

	a := 10
	b := 20

	//电脑匿名函数，参数传递过去
	defer func(a,c int){
		fmt.Printf("内部：a = %d,b = %d\n",a,b)
	}(10,20)

	a = 111
	b = 222

	fmt.Printf("外部：a = %d, b = %d\n", a, b)

}