package main

import "fmt"

func main()  {

	/**
		切片和数组的区别
		1）数组[]里面的长度是固定的一个常量，数组不能修改长度。len 和 cap 永远都是 5
		2) 切片[]里面为空，或者为 ,,,  切片的长度和容量可以不固定
	 */

	arr := [5]int{}
	fmt.Println("len(arr)=", len(arr) , "cap(arr)=", cap(arr))


	//2) 切片[]里面为空，或者为 ,,,  切片的长度和容量可以不固定
	slip := []int{}
	slip = append(slip, 11)
	slip = append(slip, 22)
	fmt.Println("slip=", slip)

	/**
		切片的创建方式
	 */

	//1、自动推导类型，同时初始化
	s1 := []int{1,2,3,4}
	fmt.Println("s1 = ", s1)

	//2、借助make函数创建
	s2 := make([]string, 5, 10)
	fmt.Println("s2 = ", s2, "len(s2)=", len(s2), "cap(s2)= ", cap(s2))

	//3、没有指定容量，则容量和长度一样
	s3 := make([]int, 10)
	fmt.Println("s3 = ", s3, "len(s3)=", len(s3), "cap(s3)= ", cap(s3))
}
