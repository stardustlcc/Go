package main

import "fmt"

//函数被调用时，x 才分配空间，才初始化为0
func test01() int  {
	var x int //没有初始化，值为0
	x++
	return x  //函数调用完毕，x自动释放
}

//函数的返回值是一个匿名函数，返回一个函数类型
func test02() func() int  {
	var x int
	return func() int{
		x++
		return  x
	}
}

func main()  {

	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())

	//返回值为一个匿名函数，返回一个函数类型，通过f1 来调用返回值的匿名函数，f1 来调用闭包函数
	//它不关心这些捕获了的变量和常量是否已经超出了作用域
	//所以只有闭包还在使用它，这些变量就还会存在
	f1 := test02()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

}
