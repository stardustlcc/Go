package main

import (
	"fmt"
	"runtime"
)

func main()  {

	go func() {
		for i := 0; i<5; i++ {
			fmt.Println("222222")
		}
	}()

	for i := 0; i<2; i++ {
		//让出时间片，让别的协程执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("1111")
	}

}
