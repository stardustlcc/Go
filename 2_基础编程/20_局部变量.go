package main

import "fmt"

func test()  {
	a := 10
	fmt.Println("a=",a)
}

func main()  {
	/**
		特点：
			定义在{}中的变量是局部变量
			执行到定义变量的那句话，才开始分配空间，离开作用域自动释放
			作用域，变量其作用的范围
	 **/

	test()
	//a = 111 error

	//
	if flag := 3; flag == 10 {
		fmt.Println("flag=", flag)
	} else {
		flag = 10
		fmt.Println("flag=", flag)
	}
	//flag = 11 error




}
