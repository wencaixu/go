package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	rand.Seed(time.Now().UnixNano())

	var seeds [10]int

	for i := 0; i < len(seeds); i ++{
		seeds[i] = rand.Intn(10)
	}

	//冒泡排序
	for i := 0; i < len(seeds); i++{
		for j := i; j < len(seeds); j++{
			if seeds[i] >= seeds[j]{
				TMP := seeds[i]
				seeds[i] = seeds[j]
				seeds[j]=TMP
			}
		}
	}

	for i,data := range seeds{
		fmt.Printf("seeds[%d]->%d\n",i,data)
	}
}