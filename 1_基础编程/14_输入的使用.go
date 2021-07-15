package main

import "fmt"

func main()  {

	var a int
	fmt.Printf("请输入变量a:")

	//阻塞等待用户的输入， 定义格式 Scanf
	//fmt.Scanf("%d", &a)

	//更方便的方式， 不需要定义格式 Scan
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
