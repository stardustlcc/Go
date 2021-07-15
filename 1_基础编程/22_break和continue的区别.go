package main

import (
	"fmt"
	"time"
)

func main()  {

	i := 0
	for { //相当于死循环
		i++

		if i == 2 {
			continue
		}

		fmt.Println("i=", i)
		time.Sleep(time.Second * 1)

		if i == 5 {
			break //跳出循环，跳出最近的一层循环
		}


	}

}
