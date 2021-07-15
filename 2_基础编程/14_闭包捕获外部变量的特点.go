package main

import "fmt"

func main()  {

	a := 100
	str := "hello"

	//闭包里面的参数可以覆盖外部变量的值。不关心作用域
	func() {
		a = 100
		str = "hello"
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}()

	fmt.Println("a=", a)
	fmt.Println("str=", str)

}
