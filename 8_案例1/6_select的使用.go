package main

import (
	"fmt"
	"time"
)

func main()  {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func( ch1 chan int) {
		ch1 <- 100
	}(ch1)

	go func(ch2 chan string) {
		ch2 <- "hello"
		time.Sleep(3 * time.Second)
	}(ch2)

	//如果程序所有 goroutine 都退出了，<-ch 这里就会阻塞，不可能等到任何东西会报错
	for {
		select {
		case  num:= <-ch1:
			fmt.Println("num=", num)
		case  msg := <-ch2:
			fmt.Println("msg=", msg)
		case <- time.After(5*time.Second): //超时没有消息
			fmt.Println("select is time out done")
			break
		}
	}

}
