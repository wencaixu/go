package main

import "fmt"

func main(){
	m := map[int]string{1:"mike",2:"vivo"}
	fmt.Println("m = ",m)
	delete(m,1) //删除key值为1
	fmt.Println("m = ",m)
}