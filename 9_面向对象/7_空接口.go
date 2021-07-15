package main

import "fmt"

type test interface {

}

/**
	空接口是没有定义任何方法的接口
	因此，任何类型都可以实现空接口
	空接口类型的变量可以存储任意类型的变量

	使用空接口实现可以接收任意类型的函数参数。
 */
func main()  {

	var t test
	str := "你好"
	t = str
	fmt.Println("t=", t)

	num := 100
	t = num
	fmt.Println("t=", t)

	slip := []int{1,2,3,4}
	t = slip
	fmt.Println("t=", t)

	var x interface{}
	x = slip
	fmt.Println("x=", x)

	testFun("你好")
	testFun(1000)
	testFun(3.14)
	testFun(slip)
}

func testFun(inter interface{})  {
	fmt.Println(inter)
}
