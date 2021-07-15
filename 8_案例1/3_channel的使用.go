package main

import "fmt"

func getMessage( ch chan string)  {
	msg := <-ch
	fmt.Println("msg = ", msg)
}

func onlyGetFun( ch <-chan int)  {
	num := <- ch
	fmt.Println("num=", num)
}

func onlySendFun( ch chan<- int, num int) {
	ch <- num
}

func main()  {

	/**
		背景:
		如果只是实现了go 函数的 并发编程，是没有特别意义的。要保证协程与协程之间的数据交换才有意义。
		如果采用共享内存来实现数据交换，会产生资源竞争的问题。为了保证交换数据的正确性，若采用了互斥量对内存加锁，便会造成性能问题。
		此时，go 采用了 channel , 即，通信实现共享内存 而不是通过共享内存来实现通信

		介绍:
		go 中的 channel 类似一个传送带 或 队列，它准循着先入先出的规则，来保证发送数据的顺序
		go 中的 channel 都是可以申请类型的，在声明channel的时候需要指定元素类型

		channel 是引用类型,默认值是nil, 因为是引用类型，所以 channel 必须通过 make 函数初始化申请内存空间之后才能使用
	 */


	//var ch1 chan int   // 声明一个传递整型的通道
	//var ch2 chan bool  // 声明一个传递布尔型的通道
	//var ch3 chan []int // 声明一个传递int切片的通道

	/**
		创建 channel 的格式
		make(chan 元素类型, [缓冲大小])
		channel 分
			1）有缓存的通道（）
			2）无缓存的通道 (阻塞通道，没有信息的时候是不能接收的，如果消息没有被读出，是不能继续发送的)

	 */

	/**
		有缓存的通道
	 */
	ch1 := make(chan int, 3)
	ch1 <- 100
	ch1 <- 111
	ch1 <- 222

	//ch1 <- 333 容量不够，报错哦

	fmt.Println("len(ch1)=", len(ch1), "cap(ch1)=", cap(ch1))
	fmt.Println("ch1=", <-ch1)

	/**
		无缓存的通道
		注意：无缓存的通道，只有在有人接收值的时候才能发送值，否则就会报错
		解决方式：启用一个goroutine去监听阻塞接收值
	 */
	ch2 := make(chan string)
	// ch2 <- "这样不好"

	//定义一个，接收值的协程
	go getMessage(ch2)

	//有接收值，我就可以推送消息了
	ch2 <- "你好，我是测试通道"
	fmt.Println("len(ch2)=", len(ch2), "cap(ch2)=", cap(ch2))



	/**
		channel 的操作
		通道有发送（send）、接收(receive）和关闭（close）三种操作
			1）发送到通道   ch <- 值
			2）从通道接收   <- ch
			3）关闭通道     close(ch)
			注意：关闭通道不是必须的，它可用被垃圾回收机制正常回收。
			对关闭的通道进行操作会发生什么呢？
				1）已关闭，在发送内容， 会 panic 错误
				2）已关闭，继续接收内容，若有内容则可以接收，直到通道为空
				3）已关闭，继续接收内容，若没有内容则得到0值
				4）已关闭，继续关闭，会 panic 错误
	 */

	close(ch2)
	close(ch1)

	//ch2 <- "你好，我又来了"  //panic: send on closed channel
	//close(ch2)  //panic: close of closed channel

	//对 channel 进行遍历，阻塞取值，直到没有值时才停止
	for val := range ch1 {
		fmt.Println("ch1=", val)
	}

	// 测试已关闭，继续接收值，会得到 0 值
	var i = 0
	for {
		i++
		fmt.Println(<-ch1)
		if i >= 5 {
			break
		}
	}


	/**
		channel 作为函数参数传递的时候，可以对 channel 做限制；比如：只能接收 or 只能发送
			1）chan<- int  是一个只能发送的通道，可以发送但是不能接收
	    	2）<-chan int  是一个只能接收的通道，可以接收但是不能发送
	 */

	ch3 := make(chan int)

	go onlyGetFun(ch3)
	onlySendFun(ch3, 100)











}
