package main

import (
	"fmt"
	"sync"
)

var waitChild sync.WaitGroup

func test(i int)  {
	defer waitChild.Done()
	fmt.Println("我是子协程 ==", i)
}

func main()  {
	/**
		启动多个goroutine时
		使用了sync.WaitGroup来实现goroutine的同步
		下面的程序每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
	 */
	for i := 0; i< 10; i++{
		waitChild.Add(1)
		go func(num int) {
			defer waitChild.Done()
			fmt.Println("我是子协程 ==", num)
		}(i)
	}

	waitChild.Wait()
}
