package main

import (
	"fmt"
	"time"
)

func newTask()  {
	for  {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}

func main()  {
	//创建一个协程
	go newTask()
	for  {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}





}
