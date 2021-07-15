package main

import (
	"fmt"
	"time"
)

func main()  {

	//创建一个无缓存的channel
	//无缓存的通道，无论是写，还是读都会阻塞
	ch := make(chan int)

	fmt.Printf("len=%d, cap=%d\n",len(ch), cap(ch))

	go func() {

		for i:=0;i<3; i++{
			fmt.Println("子协程 i =", i)
			ch <- i
		}

	}()

	time.Sleep(2* time.Second)

	for i:= 0; i<3; i++ {
		num := <- ch
		fmt.Println("num=", num)
	}


}
