package main

import "fmt"

func testa(){
	fmt.Println("aaaaaaa")
}


func testb(){
	panic("bbbbbbbb")
}

func testc(){
	fmt.Println("cccccccc")
}


func main(){
	testa()
	testb()
	testc()
}