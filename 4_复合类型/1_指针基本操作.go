package main

import "fmt"

/**
	go 默认值 nil, 没有 null 常量
	操作符 "&" 取变量地址，"*" 通过指针访问目标对象
	不支持指针运算，不支持"->"运算符，直接使用"."访问目标成员
**/
func main()  {
	/**
		每个变量有两层含义
		一个是变量的内存
		一个是变量的地址
	**/
	var a int
	a = 10

	fmt.Println("a=", a)
	fmt.Printf("&a=%v\n",&a)
	fmt.Printf("&a=%v\n",*&a)

	/**
		保存某个变量的地址，需要指针类型
		*int 保存 int的地址
		**int 保存 *int地址
	**/

	var p *int
	p = &a
	fmt.Printf("*p=%v\n", *p)

	*p = 100
	fmt.Printf("*p=%v\n", *p)


	/**
		不要操作没有指向的内存
	**/

	//var s *int
	//s = nil
	//*s = 11
	//fmt.Printf("*s=%v\n", *s)
}
