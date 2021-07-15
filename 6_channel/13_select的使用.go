package main

import "fmt"

func fibonacci(ch chan <- int, quit <-chan bool)  {
	x, y := 1, 1
	for {
		//监听channel数据流动
		select {
		case ch <- x:
			x,y = y,x+y
		case flag := <-quit:
			fmt.Println("flag=", flag)
			return
		}
	}
}

func main()  {

	/**
		select 的 case 必须是 io 操作
	 */

	ch := make(chan int) 	//数字通信
	quit := make(chan bool) //程序是否结束

	//消费者，从channel中读取内容
	go func() {
		for i:=0; i<8; i++{
			num := <-ch
			fmt.Println("num=", num)
		}
		quit<-true
	}()
	//生产者，往channel中发内容
	fibonacci(ch, quit)

}
