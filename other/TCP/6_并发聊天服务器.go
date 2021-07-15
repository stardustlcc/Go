package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //管道：用于给用户发送数据
	Name string
	Addr string
}

//保存在线用户
var onlineMap map[string]Client
var message   = make(chan string)



func WriteMsgToClient(client Client, conn net.Conn)  {
	for msg := range client.C {
		conn.Write([]byte(msg+"\n"))
	}
}

func MakeMsg(client Client, msg string) (buf string)  {
	buf = "[" + client.Addr + "]" + client.Name + ": "+ msg
	return
}

//处理用户连接
func HandleConn(conn net.Conn)  {

	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体, 默认，用户名和网络地址一样
	client := Client{make(chan string), cliAddr, cliAddr}

	//把结构体添加到map
	onlineMap[cliAddr] = client

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(client, conn)

	//广播某个人在线
	message<-MakeMsg(client, "login")

	//提示，我是谁
	client.C <- MakeMsg(client,"I am here")

	//对方是否主动退出
	isQuit := make(chan bool)

	//对方是否有数据发送
	hasData := make(chan bool)

	//接收用户发送过来的数据
	go func() {

		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)
			if n == 0 {//对方断开，或出问题
				isQuit <- true
				fmt.Println("conn.Read  err=", err)
				return
			}

			msg := string(buf[:n-2])

			if len(msg) == 3 && msg == "who" {
				//遍历当前map,给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _,tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//rename|mike
				name := strings.Split(msg,"|")[1]
				client.Name = name
				onlineMap[cliAddr] = client
				conn.Write([]byte("rename ok\n"))
			} else {
				//需要转发此内容
				message<-MakeMsg(client, msg)
			}

			hasData <- true


		}

	}()

	for {
		//通过select 检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)//当前用户从map移除
			message<-MakeMsg(client,"login out") //广播谁下线了
			return
		case <-hasData:
		case <-time.After(30*time.Second)://60秒后超时
			delete(onlineMap,cliAddr)
			message <- MakeMsg(client, "time out leave out")//广播下线了
			return
		}
	}
}

func Messager()  {

	//给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message//没有消息的时候就阻塞了
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main()  {

	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen err=", err)
	}

	defer listen.Close()

	//协程1：转发消息（有人发消息了，则给每个成员发送消息）
	go Messager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			continue
		}

		//处理用户连接
		go HandleConn(conn)
	}
}
