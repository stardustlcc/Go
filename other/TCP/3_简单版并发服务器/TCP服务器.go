package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn)  {
	//函数调用完毕，自动完毕
	defer conn.Close()

	//获得客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect successful")

	buf := make([]byte, 2048)

	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}

		fmt.Printf("[%s]:%s len %d \n",addr, string(buf[:n-1]), len(string(buf[:n])))

		if "exit" == string(buf[:n-1]) {
			fmt.Println(addr, ", exit")
			return
		}

		conn.Write([]byte( strings.ToUpper(string(buf[:n])) ))
	}

}

func main()  {

	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
	}

	defer listen.Close()

	for {
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("err1=", err)
		}


		//处理用户请求,每来一个用户，新建一个协程
		go HandleConn(conn)

	}



}
