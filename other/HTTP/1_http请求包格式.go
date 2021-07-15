package main

import (
	"fmt"
	"net"
)

func main()  {

	listen, err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen  err=", err)
		return
	}

	defer listen.Close()

	//阻塞等待用户的链接
	conn, err1:= listen.Accept()
	if err1 != nil {
		fmt.Println("listen.Accept err1=", err1)
		return
	}

	defer conn.Close()


	//接收客户端的数据
	buf := make([]byte, 1024*4)

	n, err2 := conn.Read(buf)
	if err2!=nil {
		fmt.Println("conn.Read err2=", err2)
		return
	}

	if n == 0 {
		fmt.Println("n==0")
		return
	}

	//浏览器请求：127.0.0.1：8000

	//请求行 GET / HTTP/1.1
	//请求头
	//请求包
	fmt.Printf("#%v#", string(buf[:n]))

}
