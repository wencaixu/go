package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//设置种子,如果随机种子参数一样，每次运行程序产生的随机数一样
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5;i++{

		fmt.Println(rand.Intn(100))//限制在100以内的数
	}
}
