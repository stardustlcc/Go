package main

import (
	"fmt"
	"time"
)

func main()  {

	//创建一个有缓存的channel
	//有缓存的通道，只有写满了才会阻塞，读完后才能写
	ch := make(chan int,3)

	fmt.Printf("len=%d, cap=%d\n",len(ch), cap(ch))

	go func() {

		for i:=0;i<10; i++{
			fmt.Println("子协程 i =", i)
			ch <- i
			//fmt.Printf("len=%d, cap=%d\n",len(ch), cap(ch))
		}

	}()

	time.Sleep(2* time.Second)

	for i:= 0; i<10; i++ {
		num := <- ch
		fmt.Println("num=", num)
	}


}
