package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn)  {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open=", err)
		return
	}
	defer f.Close()

	//读文件内容，读多少就给接收方发送多少
	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err = ", err)
			}
			return
		}

		//发送内容
		conn.Write(buf[:n]) //给服务器发送内容
	}


}

func main()  {
	// 提示输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)

	//获取文件名
	info, err:= os.Stat(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 !=nil {
		fmt.Println("err1=", err1)
		return
	}
	defer conn.Close()

	//给接收方，先发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err !=nil {
		fmt.Println("err write=", err)
		return
	}

	//接收对方的回复，如果回复"ok"表示对方准备好了，准备发文件
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err !=nil {
		fmt.Println("err read=", err)
		return
	}

	if "ok" == string(buf[:n]) {
		//发送文件内容
		SendFile(path, conn)
	}


}
