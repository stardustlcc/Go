package main

import (
	"fmt"
	"runtime"
)

func main()  {

	//设置cpu的核数

	n := runtime.GOMAXPROCS(2)
	fmt.Println("n=", n)// 4 核

	//指定单核
	for {
		go fmt.Print(1)

		fmt.Print(0)
	}
}
