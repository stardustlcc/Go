package main

import "fmt"

//只能写，不能读
func producter( out chan<- int )  {
	for i := 0; i<10; i++ {
		out<-i
	}
	close(out)
}

//只能读，不能写
func consumer (in <-chan int)  {
	for num := range in {
		fmt.Println("num=", num)
	}
}

func main()  {

	//创建一个双向通道
	ch := make(chan int)

	//生产者，产生数字，写入channel
	go producter(ch)

	//消费者，从channel中读取数据
	consumer(ch)

}
