package main

import "fmt"

func main()  {

	//a := 10

	var p *int
	//p = &a

	//p 是 *int, 指向int类型
	p = new(int)
	*p = 123

	fmt.Printf("*p=%v\n", *p)


	q := new(int)
	*q = 4444
	fmt.Printf("*q=%v\n", *q)


}
