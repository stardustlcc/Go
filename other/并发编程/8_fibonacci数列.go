package main

import "fmt"

//ch 只写  quit 只读
func fibonacci(ch chan<- int, quit<-chan bool)  {
	x, y := 1,1
	for {
		select {
		case ch<- x:
			x, y = y, x+y
		case flag:= <-quit:
			fmt.Println("flag=", flag)
			return
		}
	}
}

func main()  {

	ch := make(chan int) //数字通信
	quit := make(chan bool) //程序是否结束

	//消费者， 从channel中读取内容
	go func() {
		for i:=0; i<8; i++ {
			num := <-ch
			fmt.Println("num=", num)
		}
		quit<- true
	}()

	//生产者，产生数字，写入channel
	fibonacci(ch, quit)
}
