package main

import "fmt"

func main(){
	m := map[int]string{1:"mike",2:"vov"}

	for key,value := range m{
		fmt.Printf("%d--->%s",key,value)
	}

	//判断一个key值是否存在
	//第一个返回值为key对应的value,第二个返回值是key是否存在的条件，存在的话ok为true
	value,ok := m[0]
	if ok == true{
		fmt.Println("m[1] = ",value)
	}else{
		fmt.Println("key不存在")
	}

}