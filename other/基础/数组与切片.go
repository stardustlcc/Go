package 基础

import "fmt"

func main()  {

	//数组，固定长度，所以很少使用

	// 声明一个数组
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// 声明并赋值
	arr1 := [3]int{1,2,3}
	fmt.Println(arr1)

	var arr2 = [3]int{1,2,3}
	fmt.Println(arr2)

	arr3 := [...]string{"你好", "我是", "中国人", "是世界"}
	fmt.Println(arr3)
	//[你好 我是 中国人 是世界]

	fmt.Printf("arr2 %d 的类型是 %T \n", arr2, arr2)
	//arr2 [1 2 3] 的类型是 [3]int

	fmt.Printf("arr3 %s 的类型是 %T \n", arr3, arr3)
	//arr3 [你好 我是 中国人 是世界] 的类型是 [4]string

	// 定义字面量
	type arr3Int [3]int
	arr4 := arr3Int{1,2,3}
	fmt.Println(arr4)

	//通过make 声明一个切片
	s1 := make([]string, 3)
	fmt.Printf("%d \n", len(s1))//3
	fmt.Printf("%d \n", cap(s1))//3
	fmt.Printf("%s \n", s1)

	s2 := make([]int, 2, 8)
	fmt.Printf("%d\n", len(s2))
	fmt.Printf("%d\n", cap(s2))
	fmt.Printf("%d\n",s2)

	//设置一个数组
	s3 := [3]int{1,2,3}
	fmt.Printf("%d\n", len(s3))
	fmt.Printf("%d\n", cap(s3))
	fmt.Printf("%d\n", s3)

	//通过切片表达式基于数组或切片生成新切片
	s4 := []int{1,2,3,4,5,6,6}
	s5 := s4[2:4]
	fmt.Printf("%d\n", s5)
	fmt.Printf("%d\n", len(s5))
	fmt.Printf("%d\n", cap(s5))






	
}