package main

import "fmt"

func main()  {

	/**
		字符：单引号 (字符，往往都只有一个字符，除转义符'\n')
		字符串：双引号（字符串有1个或多个字符组成，都是隐藏了一个结束符'\0', \0是看不到的没有长度的）
	**/
	var ch byte
	var str string

	ch = 'a'
	str = "a"

	fmt.Println("ch=", ch)
	fmt.Println("str=", str)
	fmt.Printf("len(str) =%d\n", len(str))


	/**
		只想操作字符串中的某个字符
	**/
	str1 := "hello"
	fmt.Printf("str1[0]=%c str1[1]=%c\n", str1[0], str1[1])


}
