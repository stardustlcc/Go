package main

import (
	"fmt"
	"time"
)

func print(str string)  {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}

func p1()  {
	print("hello")
}

func p2()  {
	print("world")
}
func main()  {

	go p1()
	go p2()

	for {

	}
}
