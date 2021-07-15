package main

import "fmt"

func main()  {

	fmt.Println("4 > 3 结果：", 4>3)
	fmt.Println("4 != 3 结果：", 4!=3)

	/**
		逻辑运算符 ! && ||
	**/

	//!非
	fmt.Println("!(4>3) 结果：", !(4>3))
	fmt.Println("!(4!=3) 结果：", !(4!=3))

	//&& 与
	fmt.Println("true && true ==", true && true)
	fmt.Println("true && false ==", true && false)

	//|| 或
	fmt.Println("true || false ==", true || false)
	fmt.Println("false || false ==", false && false)

	a := 8
	//fmt.Println("0<=a <= 10", 0<=a<= 10) // 不兼容，err
	fmt.Println("(0<=a && a <= 10) == ", 0<=a && a<= 10)
	fmt.Println("(0<=a && a <= 2) == ", 0<=a && a<= 2)

}
