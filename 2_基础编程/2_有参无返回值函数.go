package main

import "fmt"

func MyFuncParam1(num int) {
	num = 100
	fmt.Println("num=", num)
}

func MyFuncParam2(num int) {
	fmt.Println("num=", num)
}

func MyFuncParam3(num, age int, score float64) {
	fmt.Println("num=", num)
	fmt.Println("age=", age)
	fmt.Println("score", score)
}

func main()  {
	var num int
	MyFuncParam1(num)
	MyFuncParam2(666)
	MyFuncParam3(999, 18, 3.14)
}
