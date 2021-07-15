package main

import "fmt"

func Add(a,b int) int  {
	return  a + b
}

func Minus(a,b int) int  {
	return  a - b
}

//函数也是一种数据类型，通过 Type 给一个函数类型起名
//FuncType 代表的是一个函数类型，有两个参数，一个返回值
type FuncType func(int,int) int

func main()  {

	sum := Add(100,200)
	fmt.Println("sum=", sum)

	dis := Minus(200, 100)
	fmt.Println("dis=", dis)


	//定义函数变量
	var Ftest FuncType

	Ftest = Add //是变量就可以赋值

	count := Ftest(10, 20)
	fmt.Println("count=", count)

	Ftest = Minus
	min := Ftest(20,10)
	fmt.Println("min=", min)

}
