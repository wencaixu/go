package main

import "fmt"

func modify(a [5]int){
	a[0]=666
	fmt.Println("modify a = ",a)
}


func modifyT(a *int){
	a[0]=666
	fmt.Println("modify a = ",a)
}

func main(){
	a := [5]int{1,2,3,4,5} //初始化

	//modify(a) //传递数组参数

	//modifyT(&a)

	fmt.Println("main:a=",a)
}