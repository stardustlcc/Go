package main

import "fmt"

func main()  {

	//传统写法，变量的定义声明
	var a int
	var b float64

	a, b = 10, 3.14
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//批量处理 不同类型变量的声明
	var (
		name string
		sex = 1 //自动推导类型
		age int = 18
		sort float64
	)

	name = "KK"
	sort = 3.14

	fmt.Println("name=", name)
	fmt.Println("sex=", sex)
	fmt.Println("age=", age)
	fmt.Println("sort=", sort)


	//批量处理 不同类型常量的声明
	const (
		number = 10 //自动推导类型
		float float64 = 3.14
	)

	fmt.Println("number=", number)
	fmt.Println("float=", float)

}
