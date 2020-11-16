package main

import "fmt"

func main(){
	a:=10
	str := "mike"
	func(){
		a = 666
		str = "go"
		fmt.Println(a, " " , str)
	}()
	fmt.Println(a," ",str)
}