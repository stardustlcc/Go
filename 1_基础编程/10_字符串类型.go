package main

import "fmt"

func main()  {

	//声明变量,字符串是双引号
	var str1 string
	str1 = "abc"
	fmt.Println("str1 =", str1)

	//自动推导类型
	str2 := "hello!"
	fmt.Printf("str2=%s type=%T\n", str2,str2)

	//内建函数，len()可以测字符串的长度，有多少个字符
	fmt.Println("len = ", len(str2))
}
