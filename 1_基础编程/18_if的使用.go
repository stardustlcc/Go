package main

import "fmt"

func main()  {

	str := "aa"

	//没有圆括号，花括号要在条件的同一行
	if str == "aa" {
		fmt.Println("is ok")
	} else {
		fmt.Println("is error")
	}

	//if 支持 1 个初始化语句, 初始化语句和判断条件以分号分隔
	if a := 10; a == 10 {
		fmt.Println("is ok")
	}

}
