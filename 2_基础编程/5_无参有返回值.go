package main

import "fmt"

//有返回值,必须要加return
func myFun1() int {
	a := 666
	return a
}

// 有返回值名称时，可以不用返回 result,直接 return 即可
func myFun2() (result int) {
	result = 777
	return
}

func main()  {

	a := myFun1()
	fmt.Println("a=", a)

	b := myFun2()
	fmt.Println("b=", b)
}
