package main

import "fmt"

func main()  {

	//1、声明变量
	var a bool
	a = true
	fmt.Println("a = ", a)

	//自动推导类型
	var b = false
	fmt.Printf("b=%t type=%T\n", b,b)

	c := false
	fmt.Printf("c=%t type=%T\n", c,c)
}
