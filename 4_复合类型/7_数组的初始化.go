package main

import "fmt"

func main()  {

	//声明定义同时赋值，叫初始化
	//全部初始化
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println("a=", a)

	//部分初始化，没有初始化的字符串为""，整型默认值为0
	b := [5]string{"aa","bb","cc"}
	fmt.Println("b=", b)

	//指定某个元素初始化
	c := [5]int{2:100, 4:100}
	fmt.Println("c=", c)



}
