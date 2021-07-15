package main

import "fmt"

func main()  {
	/**
		数组是指一系列同一类型数据的集合，数组中包含的每个数据被称为数组元素
		一个数组包含的元素个数被称为数组的长度
		数组长度必须是常量，且是类型的组成部分
		[2]int 和 [3]int 是不同类型
	**/

	var a [10]int
	var b [5]string
	var c [2]float64

	a[0] = 100
	b[0] = "hai"
	c[0] = 3.14

	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)


	//定义数组时，指定的数组元素个数必须是常量
	//num := 10
	//var arr [num]int error

	/**
		操作数据元素，从0开始，到 len()-1
		这里的下标，可以是变量或常量
	 */
	var arr [10]int
	arr[0] = 100
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	fmt.Println("arr=", arr)

}
