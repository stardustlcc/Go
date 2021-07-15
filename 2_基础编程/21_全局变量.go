package main

import "fmt"

//定义在函数外部的变量，叫全局变量
//全局变量在任何地方都能使用
var a int = 10

func test()  {
	fmt.Println("test a =", a)
}

func main()  {

	a = 100
	fmt.Println("a = ", a)

	test()

}
