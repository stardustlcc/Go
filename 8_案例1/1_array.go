package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan string)

func main()  {

	go func() {
		chan1 <- "hello"
	}()

	for {
		select {
		case <- chan1:
			fmt.Println("chan1")
		case <- time.After(10* time.Second):
			fmt.Println("time out !!!")
		}
	}

}
