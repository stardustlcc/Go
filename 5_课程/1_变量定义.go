package main

import "fmt"

//函数外定义全局变量
var aa = 3
var bb = "string is this"

var (
	cc int
	dd string
	ee bool
)

func main()  {

	//1、定义变量不赋值
	var a int
	var s string
	fmt.Println(a, s)

	//2、变量赋初值
	var num, age int = 100,18
	var str string = "abc"
	fmt.Println(num,age,str)

	//3、自动推断类型
	var name,sex,desc = "cici", 18, false
	fmt.Println(name, sex, desc)

	//4、简短的定义方式
	str1 := "你看看，我这是字符串"
	fmt.Println(str1)

	//5、函数外定义变量
	fmt.Println(aa,bb,cc,dd,ee)

}
