package main

import (
	"fmt"
	"os"
)

func main()  {

	//接收用户命令行参数，都是以字符串方式传递
	list := os.Args
	fmt.Println("len(list) = ", len(list))

	for k, v := range list {
		fmt.Printf("list[%d] - %s\n", k,v)
	}

}
