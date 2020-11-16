package main

import "fmt"

func main(){

	//定义数组
	var id [10]int
	var bs [11]int

	fmt.Printf("id %d,bs %d",len(id),len(bs))

	//定义数组，制定的数组元素个数是常量

	id[2] = 1

	for i,date := range id{
		fmt.Printf("a[%d] = %d\n", i, date)
	}
}