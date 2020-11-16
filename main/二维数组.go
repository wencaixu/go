package main

import "fmt"

func main(){

	var id [3][4]int
	k := 0

	for i := 0; i < 3; i++{
		for j:=0;j < 4; j++{
			k++
			id[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, id[i][j])
		}
		fmt.Println()
	}

	fmt.Println("a=",id)

	fmt.Println()

	//初始化二维数组
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8},{9, 10, 11, 12}}
	fmt.Println("b = ", b)
	//部分初始化，没有初始化的值为0
	c := [3][4]int{{1,2,3},{5,6,7,8},{9,10}}
	fmt.Println("c = ", c)
	d := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println("d = ", d)

	e := [3][4]int{1: {5, 6, 7, 8}}
	fmt.Println("e = ", e)
}