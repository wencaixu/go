package main

import "fmt"

func main(){

	var p *int

	p = new(int)

	//p是*int类型，指向int类型
	*p = 666

	fmt.Printf("*p= %d",*p)

	q := new(int) //自动推导类型

    *q = 777
    fmt.Println("*q = ", *q)
}