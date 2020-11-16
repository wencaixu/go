package main

import "fmt"

type People struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main(){
	var p = People{1,"Jetty",'m',12,"中国"}
	fmt.Println("p = ",p)

	s2 := People{name:"mike",addr:"bj"}
	fmt.Println("s2 = ",s2)
}