package main

import (
	"fmt"
)

func main()  {

	isOk := make(chan bool)

	/**
		go 语言中的协程寿命取决于 go 的主协程。
		主协程执行完毕后，go 的子协程会一并销毁。
		因为，go 的子协程的创建需要花费一些时间。所以如果go 的主协程执行完了，子协程也就没有执行的机会了。
	 */
	go func() {
		fmt.Println("我是子协程")
		isOk <- true
	}()

	fmt.Println("hello is main go")


	//1）加一个死循环，等待子协程执行完
	//for {
	//
	//}

	//2）加一个睡眠时间
	//time.Sleep(3*time.Second)

	//3) 使用通道交互，如果通过没有消息，会阻塞go主协程
	<- isOk



}
