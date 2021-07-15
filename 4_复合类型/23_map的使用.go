package main

import "fmt"

func main()  {

	/**
		go 语言中的map（字典，映射）是一种内置的数据结构，它是一个"无序"的key->value对的集合
		在一个map里面所有的键都是唯一的，而且必须是支持== 和 ！= 操作符的类型
		切片，函数以及包含切片的结构体由于具有引用语义，不能作为映射的键

		map 只有len 没有 cap
	 */


	//1、定义一个map 变量
	var map1 map[int]string
	fmt.Println("map1 =", len(map1))

	//2、通过make 创建
	map2 := make(map[int]string)
	fmt.Println("map2=", len(map2))

	//3、通过make 创建，可以指定长度，长度也是自动扩容的
	map3 := make(map[int]string,2)
	fmt.Println("map3=",len(map3))//0

	map3[1] = "aaa"
	map3[2] = "bbb"
	map3[3] = "ccc"
	fmt.Println("map3=", map3)

	map4 := map[int]string{1:"aaa", 2:"bbbbb", 5:"cccc"}
	fmt.Println("map4=", map4)

}
