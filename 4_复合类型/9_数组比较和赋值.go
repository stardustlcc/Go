package main

import "fmt"

func main()  {

	/**
		数组支持比较
		只支持 == 和 !=
		比较是不是每一个元素都一样，2个数组比较，数组类型要一样
	 */
	a := [3]int{1,2,3}
	b := [3]int{1,200,300}
	c := [4]string{"a","b"}
	d := [4]string{"a","b"}

	fmt.Println("a==b", a==b)
	fmt.Println("d==c", d==c)



}
