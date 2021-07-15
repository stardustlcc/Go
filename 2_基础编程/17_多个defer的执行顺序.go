package main

import (
	"fmt"
)

func test(t int) int {
	return 100 / t
}

func main()  {

	//defer 是先进后出
	defer fmt.Println("aaaaaa")
	defer fmt.Println("bbbbbb")
	defer fmt.Println("cccccc")

	test(0)

	defer fmt.Println("dddddd")

}
