package main

import "fmt"

func main()  {

	//声明变量
	var f1 float32
	f1 = 3.14
	fmt.Printf("%f type =%T\n", f1, f1)

	//自动推导类型 float64
	var f2 = 3.14
	fmt.Printf("%f type =%T\n", f2, f2)

	/**
		float64 存储小数比 float32 更准确，精度更高
	**/
}
