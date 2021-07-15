package main

import "fmt"

func main()  {

	//多重赋值
	a,b,c := 1,2,3

	fmt.Println("a=",a)
	fmt.Println("b=",b)
	fmt.Println("c=",c)

	fmt.Println("-------------")

	//交换两个变量的值
	a,b = b,a
	fmt.Println("a=",a)
	fmt.Println("b=",b)

	fmt.Println("-------------")

	// _ 匿名变量的使用，丢弃数据，不处理(配合函数使用，比较有意义)
	a1, _ := a,b
	fmt.Println("a1=",a1)

	// _ 匿名变量的使用，配合函数
	var age int
	age, _ = test()
	fmt.Println("age=",age)

}

func test()(age,sex int)  {
	return 18,1
}
