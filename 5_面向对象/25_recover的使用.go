package main

import "fmt"

func TestA()  {
	fmt.Println("aaaaaaaa")
}
func TestB(x int)  {

	//设置decover
	defer func() {
		//recover()//可以打印panic的错误信息
		//fmt.Println(recover())
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()

	var a [10]int
	a[x] = 100
}
func TestC()  {
	fmt.Println("cccccccc")
}

func main()  {

	TestA()
	TestB(20)
	TestC()
}
