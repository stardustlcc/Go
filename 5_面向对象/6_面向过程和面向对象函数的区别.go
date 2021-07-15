package main

import "fmt"

//面向过程，实现2数相加
func add(a , b int) int {
	return a+b
}

//面向对象，方法：给某个类型绑定一个函数
type long int

func (tmp long) add2 ( other long) long  {
	return  tmp + other
}

func main()  {


	res := add(1,2)
	fmt.Println("a+b=", res)

	//定义一个变量
	var tmp long
	tmp = 10
	//调用方式，变量名.函数（所需参数）
	res2 := tmp.add2(5)
	fmt.Println(res2)



}
