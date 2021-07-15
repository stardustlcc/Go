package main

import (
	"fmt"
	"time"
)

func main()  {

	go func(){
		i := 0
		for  {
			i++
			fmt.Println("这是一个子协程 i=", i)
			time.Sleep(time.Second)
		}
	}()

	//主协程退出了，其他子协程也要跟着退出
	i := 0

	for  {
		i++
		fmt.Println("main i=", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}

}
