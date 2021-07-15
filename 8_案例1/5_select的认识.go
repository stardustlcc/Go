package main

import "fmt"

func main()  {

	/**
		通常情况下，我们测试一个通道，会写如下代码
	 */

	// 创建一个无缓存通道
	ch1 := make(chan int)

	//需要开启一个协程监听阻塞通道等待接收消息，如果没有这一步会报错，因为无缓存通道必须，有等待接收者
	go func(ch1 chan int) {
		for val := range ch1 {
			fmt.Println("val=", val)
		}
	}(ch1)

	//在有接收者的情况下，想通道发送消息
	ch1 <- 100

	// 必须来一个死循环，如果主协程生命周期到了，子协程也会跟着销毁，由于，创建子协程是需要花费时间的，所以需要阻塞主协程
	for {

	}


	/**
		问题： 如果开启了多个通道，都需要接收信息，就需要开启多个子协程阻塞等待
		性能不好
		go 内置了一个 select 关键字，可以同时响应多个通道的操作。
		select
			select 的语法类似 switch 有 case 和默认值。每个 case 对应一个通道的通信（接收 or 发送）select 会一直阻塞等待
			语法：
			select {
			case msg := <-ch1:
				fmt.Println("msg=", msg)
			case ch2 <- "hello":
				fmt.Println(" write hello")
			default:
				fmt.Println("this is default msg")
			}
	 */



}
