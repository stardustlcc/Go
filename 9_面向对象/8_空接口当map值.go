package main

import "fmt"

/**
	空接口可以存储任意类型的值
 */
func main()  {

	map1 := make(map[int]interface{})
	map1[0] = 10001
	map1[1] = "你好"
	map1[2] = true
	map1[3] = 3.14
	map1[4] = 'a'

	fmt.Println("map=", map1)

	//空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？
	//类型断言

}
