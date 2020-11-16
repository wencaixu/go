package main

import (
	"fmt"
	"os"
)

func main(){

	//接收用户传递的参数，以字符串方式传递
	list := os.Args

	n := len(list)

	//fmt.Print("n = ", list)

	for i := 0;i<n;i++{
		fmt.Println(list[i])
	}

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}


}
