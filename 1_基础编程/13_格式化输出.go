package main

import "fmt"

func main()  {

	a := 10
	b := "abc"
	c := 'a'
	d := 3.14

	//%T 操作变量所属类型
	fmt.Printf("%T,%T,%T,%T\n", a,b,c,d)

	//%v 自动匹配格式输出(对于字符类型不是很智能)
	fmt.Printf("%v %v %v %v\n", a,b,c,d)

	/**
		%d 整数
		%s 字符
		%c 字符
		%f 浮点类型
	**/
}
