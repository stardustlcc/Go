package main

import (
	"fmt"
)

func main()  {
	var num int
	fmt.Printf("请输入num:")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println(1)
		//break //默认就包含，可以不写
		fallthrough //不跳出循环，后面无条件执行
	case 2:
		fmt.Println(2)
		//break
		fallthrough
	case 3:
		fmt.Println(3)
		//break
	default:
		fmt.Println("default")
	}


	//支持 1 个初始化语句， 初始化语句和变量本身，以分号分隔
	switch age := 18; age {
	case 10:
		fmt.Println(10)
	case 15:
		fmt.Println(15)
	case 18:
		fmt.Println(18)
	default:
		fmt.Println("default")
	}

	//可以没有条件
	score := 85
	switch  {
	case score > 90:
		fmt.Println(">90")
	case score > 80:
		fmt.Println(">80")
	case score > 70:
		fmt.Println(">70")
	default:
		fmt.Println("0")
	}
	
	//条件可以有多个
	switch count := 4; count  {
	case 1:
		fmt.Println("1")
	case 2,3,4:
		fmt.Println("2,3,4")
	case 5,6:
		fmt.Println("5,6")
	default:
		fmt.Println("0")
	}
}
