package main

import "fmt"

func main()  {

	/**
		1、变量的声明方式 var 变量名 变量类型
		2、只是声明，并没有初始化，默认值为0
		3、在同一个{}内声明的变量唯一
	**/

	var name string
	name = "你叫什么名字？"
	fmt.Println("name=", name)

	/**
		4、同时声明多个变量
		var 变量名,变量名 变量类型
	**/

	var age, sex int
	age = 18
	sex = 1
	fmt.Println("age= ", age)
	fmt.Println("sex= ", sex)

	/**
		5、声明并赋值
	**/
	var apple string = "苹果"
	fmt.Println("apple=", apple)

	/**
		6、自动推导类型
		orange := "橙子"
	**/

	orange := "橙子"
	number := 100
	pi 	   := 3.14
	fmt.Printf("orange=%s type=%T\n", orange, orange)
	fmt.Printf("number=%d type=%T\n", number, number)
	fmt.Printf("sort=%f type=%T\n", pi, pi)
}
