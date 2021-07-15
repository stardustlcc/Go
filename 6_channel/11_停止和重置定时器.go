package main

import (
	"fmt"
	"time"
)

func main()  {
	timer := time.NewTimer(13*time.Second)
	timer.Reset(1* time.Second)
	<-timer.C
	fmt.Println("时间到")
}

func main01()  {

	timer := time.NewTimer(3*time.Second)

	go func() {
		<- timer.C
		fmt.Println("子协程可以打印了，定时器时间到")
	}()

	//停止定时器
	timer.Stop()

	for {

	}

}
