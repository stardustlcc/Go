package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件内容
func RecvFile(fileName string, conn net.Conn)()  {

	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err ", err)
		return
	}

	buf := make([]byte, 1024*4)
	//接收多少，写多少，一点不差
	for {
		n, err := conn.Read(buf)//接收对方发过来的文件内容
		if err != nil {
			if  err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err=", err)
			}
			return
		}

		if n == 0 {
			fmt.Println("n == 0 文件接收完毕")
			break
		}

		//往文件写入内容
		f.Write(buf[:n])
	}

}

func main()  {

	//监听
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listern err=", err)
		return
	}
	defer listen.Close()

	conn, err1 := listen.Accept()
	if err1 != nil {
		fmt.Println("listen.Accept err1=", err1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf) //读取对方发送的文件名称
	if err2 != nil {
		fmt.Println("conn.Read err1=", err2)
		return
	}

	fileName := string(buf[:n])

	//回复ok
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(fileName, conn)
}
