package main

import "fmt"

func main()  {

	a := 10
	str := "cici"

	//匿名函数，没有名字的函数
	f1 := func () {
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}

	f1()

	//给一个函数类型起别名
	type FuncType func()

	var f2 FuncType
	f2 = f1
	f2()


	//定义匿名函数，同时调用
	func(){
		fmt.Println("hello")
	}() //代表执行，匿名函数


	//带参数的匿名函数
	func(i,j int) {
		fmt.Println(i + j)
	}(1,2)


	//匿名函数，有参有返回值
	result := func(a,b int) (result int) {
		result = a + b
		return
	}(10,5)

	fmt.Println("result = ", result)
}
