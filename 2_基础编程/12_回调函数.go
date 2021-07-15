package main

import "fmt"

//回调函数，函数的参数是函数类型，这个函数就是回调函数
type FuncType1 func(a,b int) int

//多态，多种形态，调用同一个接口，有不同的表现，如：加减乘除
func Calc(a, b int, fTest FuncType1) (result int) {
	fmt.Println("Calc")
	result = fTest(a,b)
	return
}

func MyFunNameAdd(a,b int) (result int)  {
	result = a + b
	return
}

func MyFunNameDis(a,b int)(result int)  {
	result = a - b
	return
}

func MyFunNameMul(a,b int)(result int)  {
	result = a * b
	return
}

func main()  {

	sum := Calc(1,2, MyFunNameAdd)
	fmt.Println("sum=", sum)

	dis := Calc(100,2, MyFunNameDis)
	fmt.Println("dis=", dis)

	mul := Calc(100,2, MyFunNameMul)
	fmt.Println("mul=", mul)
}
