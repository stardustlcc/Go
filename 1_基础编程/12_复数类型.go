package main

import "fmt"

func main()  {

	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t=",t)

	//自动推导类型
	t2 := 3.2 + 1.23i
	fmt.Printf("%T\n", t2)

	//通过内建函数，取实部 和 虚部
	fmt.Println("real(t2) = ", real(t2))
	fmt.Println("imag(t2) = ", imag(t2))
}
