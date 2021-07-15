package main

import (
	"fmt"
	"time"
)

func main()  {

	//来不及执行
	go func() {
		i := 0
		for  {
			i++
			fmt.Println("1111")
			time.Sleep(time.Second)
		}
	}()

}
