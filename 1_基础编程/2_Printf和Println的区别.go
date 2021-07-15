package main

import "fmt"

func main()  {

	//一段一段处理，自动加换行
	fmt.Println("你好，我会自动换行")

	//格式化输出，不能自动换行
	fmt.Printf("a=%d b= %d c= %d", 111,222,333)

	//格式化输出，不能自动换行,加上 \n
	fmt.Printf("a=%d\n b= %d\n c= %d\n", 111,222,333)
}
