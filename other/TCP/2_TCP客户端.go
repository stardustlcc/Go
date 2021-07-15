package main

import (
	"fmt"
	"net"
)

func main()  {
	//主动链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
	}

	defer conn.Close()

	//发送数据
	conn.Write([]byte("how a you?"))
}
