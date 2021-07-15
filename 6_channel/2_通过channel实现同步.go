package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func print(str string)  {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}

//p1执行完后，才执行其他的
func p1()  {
	print("hello")
	ch <- 1 //给管道写数据
}

func p2()  {
	<-ch  //从管道取数据，如果通道没有数据就会阻塞
	print("world")
}
func main()  {

	go p1()
	go p2()

	for {

	}
}
