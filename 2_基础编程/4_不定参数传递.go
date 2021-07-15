package main

import "fmt"

//所有参数都传递给另外一个函数
func Test(args ...int)  {
	myTest(args...)

	//只想将后面两个参数传递给另外一个函数使用
	myTest2(args[len(args)-2:]...)
}

//接收不定参数
func myTest(tmp ...int)  {
	for _,data := range tmp {
		fmt.Println("myTest data=", data)
	}
}

func myTest2(tmp ...int)  {
	for _,data := range tmp {
		fmt.Println("myTest2 data=", data)
	}
}

func main()  {

	Test(1,2,3,4)

}
