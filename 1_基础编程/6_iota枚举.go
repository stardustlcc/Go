package main

import "fmt"

func main()  {

	/**
		1、iota常量自动生成器，每个行，自动累计1
		2、iota给常量赋值使用
		3、可以只写一个iota,也可以都写
	**/
	const (
		a = iota
		b = iota
		c
		d
		e
	)

	fmt.Printf("a=%d b=%d c=%d d=%d e=%d\n", a,b,c,d,e)

	/**
		4、iota 遇到const 重置为 0
	**/
	const i = iota
	fmt.Printf("i=%d\n", i)


	/**
		5、如果是同一行，值都一样
	**/
	const (
		j = iota
		k,z,x = iota,iota,iota
		o = iota
	)
	fmt.Printf("j=%d k=%d z=%d x=%d o=%d\n", j,k,z,x,o)
}
