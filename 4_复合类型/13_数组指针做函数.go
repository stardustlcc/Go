package main

import "fmt"

/**
	p 指向数组 a的地址, 它是数组指针
	*p 代表指针所指向的内存，就是实参 a
 */
func modify(p *[5]int)  {
	p[0] = 100
	fmt.Println("p=", *p)
}

func main()  {
	a := [5]int{1,2,3,4,5}

	modify(&a)

	fmt.Println("a=", a)
}