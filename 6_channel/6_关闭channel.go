package main

import "fmt"

func main()  {
	ch := make(chan int)
	go func() {
		for i:=0; i<5; i++ {
			ch <- i
		}
		//不需要写数据时，关闭channel
		close(ch)
	}()

	for num := range ch {
		fmt.Println("num=", num)
	}
}

func main01()  {

	ch := make(chan int)
	go func() {
		for i:=0; i<5; i++ {
			ch <- i
		}
		//不需要写数据时，关闭channel
		close(ch)
	}()

	for {
		//ok == true 说明管道没有关闭,为false说明已关闭
		if num, ok := <-ch ; ok == true {
			fmt.Println("num=", num)
		} else {
			fmt.Println("通道已关闭")
			break
		}
	}

}
