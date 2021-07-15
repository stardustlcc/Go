package main

import "fmt"

func TestA()  {
	fmt.Println("aaaaaaaa")
}
func TestB()  {
	fmt.Println("bbbbbbbb")
	//显示调用 panic 导致函数程序中断
	panic("this is a panic")
}
func TestC()  {
	fmt.Println("cccccccc")
}

func main()  {

	TestA()
	TestB()
	TestC()

}
