package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("cccccc")
	//return //终止此函数
	runtime.Goexit()//终止所在的协程
	fmt.Println("ddddd")
}

func main()  {

	go func() {
		fmt.Println("aaaaaa")
		test()
		fmt.Println("bbbbbb")
	}()

	//特地写一个死循环目的不让主协程结束
	for {
		//fmt.Println("")
	}
}
