package main

import "fmt"

/**
	数组是指一系列同一类型数据的集合，数组中包含的每个数据被称为数组元素
	一个数组包含的元素个数被称为数组的长度
	数组长度必须是常量，且是类型的组成部分
	[2]int 和 [3]int 是不同类型
**/
func main()  {

	var arr [50]int

	for i:=0; i<len(arr);i++  {
		arr[i] = i+1
	}
	fmt.Println("arr=", arr)
}
