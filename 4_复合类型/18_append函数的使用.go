package main

import "fmt"

func main()  {

	a := []int{}
	fmt.Println("a=", a)
	fmt.Println("len(a)=", len(a))
	fmt.Println("cap(a)=", cap(a))

	/**
		在原切片的末尾添加元素
	 */
	a = append(a,1)
	a = append(a,2)
	a = append(a,3)
	fmt.Println("a=", a)
	fmt.Println("len(a)=", len(a))
	fmt.Println("cap(a)=", cap(a))

	a1 := []int{1,2,3}
	a1 = append(a1,1111)
	a1 = append(a1,2222)
	a1 = append(a1,3333)
	fmt.Println("a1=", a1) // [1 2 3 1111 2222 3333]
	fmt.Println("len(a1)=", len(a1))
	fmt.Println("cap(a1)=", cap(a1))

}
