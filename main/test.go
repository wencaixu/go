package main

import (
	"fmt"
	"time"
)

func fun(){
	var i,sum int
	for i = 0;i <= 10000000;i++{
		sum += i
	}
}

func main(){
	start := time.Now()
	fun()
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
}

