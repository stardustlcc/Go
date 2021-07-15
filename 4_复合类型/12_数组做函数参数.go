package main

import "fmt"

/**
	数组做函数参数，它是值传递
	实参数组的每个元素给形参数组拷贝一份
	形参的数组是实参数组的复制品
 */
func modify(a [5]int)  {
	a[0] = 100
	fmt.Println("a=", a)
}

func main()  {
	a := [5]int{1,2,3,4,5}

	modify(a)

	fmt.Println("a=", a)
}
