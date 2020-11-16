package main

import "fmt"

func main(){

	var id [50] int

	//操作数组，通过下标，从0开始，到len()-1结束
	for i :=0 ; i < len(id); i++{
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}

}