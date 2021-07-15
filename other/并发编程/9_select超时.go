package main

import (
	"fmt"
	"time"
)

func main()  {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num=", num)
			case <-time.After(3*time.Second):
				fmt.Println("超时....")
				quit <- true
			}
		}
	}()

	for i:=0; i<5; i++ {
		ch <- i
		time.Sleep(3)
	}

	<-quit
	fmt.Println("程序结束。。。。")

}
