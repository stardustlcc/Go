package main

import (
	"fmt"
	"go-study/mimi"
)

func main()  {

	test002()
	fmt.Println("hai")

	res := test002()
	fmt.Println("res=", res)

	mimi.TestServices()
}
