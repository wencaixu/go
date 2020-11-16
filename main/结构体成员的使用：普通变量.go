package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main(){

	var s Student

	s.age = 10

	fmt.Println(s)

}