package main

import "fmt"

func main()  {

	//常量不允许修改值
	const a int = 10
	//a = 20
	fmt.Println("a=", a)

	//自动推导类型
	const b = 10
	fmt.Printf("b=%d type=%T\n", b, b)

	//常量不可用使用 :=
	//const c := 3.14

}
