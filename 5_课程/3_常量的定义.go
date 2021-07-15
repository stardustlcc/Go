package main

import "fmt"

func main()  {

	const filename string = "abcd.txt"
	const company = "doweidu"
	const a,b = 1111,222

	const (
		aa int = 123
		bb float64 = 3.14
		cc string = "str is this"
		dd bool = false

		aa1 = 111
		bb1 = 3.14
		cc1 = "str is this"
		dd1 = false
	)

	fmt.Printf("%T %d\n", aa,aa)
	fmt.Printf("%T %f\n", bb,bb)
	fmt.Printf("%T %s\n", cc,cc)
	fmt.Printf("%T %t\n", dd,dd)

	fmt.Printf("%T %d\n", aa1,aa1)
	fmt.Printf("%T %f\n", bb1,bb1)
	fmt.Printf("%T %s\n", cc1,cc1)
	fmt.Printf("%T %t\n", dd1,dd1)

	//常量不允许修改值
	const a1 int = 10
	//a = 20
	fmt.Println("a=", a1)

	//自动推导类型
	const b1 = 10
	fmt.Printf("b=%d type=%T\n", b1, b1)

	//常量不可用使用 :=
	//const c := 3.14

}