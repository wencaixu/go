package main

import (
	"fmt"
)

//函数基本格式说明
//func 关键字 + 函数名称 + 形参列表 + 返回值
func BasicFunction(a,b,c int,d int)(result int){
	result = a + b + c + d
	return result
}

//无参数无返回值函数
//调用方式：函数名+()
func Function(){
	a:=666
	fmt.Printf("%d",a)
}

//有参数无返回值，不定参数
//首字母小写为private 首字母大写为public
//func ParameterFunction(a int,b int)
//func ParameterFunction(a,b int)
func ParameterFunction(a ...int){
	fmt.Println(a)
}

//不定参数传递
func ParameterTransfer(a ...int){
	ParameterFunction(a...)
}

//不定参数类型
//func ParameterFunction(a int,b ...int)
func ChangeableParameterFunction(args ...int){
	for i,n := range args{
		fmt.Printf("%d %d\n",i,n)
	}
}

//含有单个返回值的函数
//方式1：func ReturnParameter()(value int)
//方式2：func ReturnParameter()(value int)
func ReturnParameter0() int {
	return 10
}

//官方推荐
func ReturnParameter1() (value int){
	value = 100
	return value
}

func ReturnParameter2()(value int){
	value = 100
	return
}

//多个返回值函数
func MultiReturnFunction()(int,string){
	a := 100
	s := "123"
	return a,s
}

//有参数有返回值
func ReturnAndParameter(a int,b int)(result int){
	result = a + b
	return result
}

//输出最大值和最小值
func ReturnMaxAndMin(a,b int)(max,min int){
	if a > b {
		max = a
		min = b
	}else if a <= b {
		max = b
		min = a
	}
	return max,min
}

//参数列表截取

//操作字符串
func GetFun1(s string)(result string){
	return s[1:]
}

//操作数组
func GetFun2(arr[] int)(result[] int){
	return arr[1:]
}

//操作不定参数
func GetFun3(s ...int)(result[] int){
    return s[1:]
}

//递归函数
func DiGui(num int)(result int){
	if num == 1{
		return 1
	}

	return num+DiGui(num - 1)
}

//函数类型
type FunctionType func(int,int) int //进行函数声明

func Add(a,b int)(result int){
	return a + b
}

func FunctionTypeModel (a int,funcType FunctionType) int{
	i := funcType(1, 2)
	return i + a
}

func BinarySearch(list []int,a int,low int,high int)(R int){
	if low <= high{
		mid := (low + high) / 2
		if list[mid] == a{
			R = mid
		}else if list[mid] > a{
			high = mid - 1
			R = BinarySearch(list,a,low,high)
		}else{
			low = mid + 1
			R = BinarySearch(list,a,low,high)
		}
	}
	return R
}

func main(){
	var list = []int{1,2,3,4,5}

	value := BinarySearch(list,4,0,4)
	fmt.Println(value)

	//函数基本格式
	fmt.Println(BasicFunction(1,2,3,4))
	//无参无返回值函数定义
	Function()
	//含有参数，但是不包含返回值
	ParameterFunction('a','b')
	//不定参数
	ChangeableParameterFunction(1,2,3)

	//不定参数传递
	fmt.Println("不定参数传递")
	ParameterTransfer(1,2,3)

	//单个返回值
	fmt.Println(ReturnParameter0())
	fmt.Println(ReturnParameter1())
	fmt.Println(ReturnParameter2())

	//多返回值
	fmt.Println(MultiReturnFunction())

	//有参数有返回值
	fmt.Println(ReturnAndParameter(1,2))

	//获取最大值和最小值
	fmt.Println(ReturnMaxAndMin(1,3))

	//操作[1:]
	fmt.Println(GetFun1("123"))
	var A = []int{1,3,2}
	fmt.Println(GetFun2(A))
	fmt.Println(GetFun3(1,2,3))

	//递归函数
	fmt.Println(DiGui(10))

	//函数类型
	fmt.Println(FunctionTypeModel(1,Add))


	//匿名函数与闭包
	f1 := func(){
		fmt.Println("123")
	}
	f1()//调用匿名函数




}
