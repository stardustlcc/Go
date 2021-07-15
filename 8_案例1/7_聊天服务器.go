package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//存储客户端的信息
type Client struct {
	C chan string //存储用户消息的通道
	Name string
	Adder string
}

//存储在线客户端的列表
var OnlineMap = make(map[string]Client)

//发送消息的通道
var message = make(chan string)

//是否还发消息
var hasData = make(chan bool)

//是否退出
var isQuit = make(chan bool)


func HanerConn( conn net.Conn)  {
	defer conn.Close()

	//获取用户的IP地址
	adderId := conn.RemoteAddr().String()
	fmt.Println("adderId=", adderId)

	/**
		1）处理数据，存储在线用户信息
		2）给当前客户端发送消息
		3）并保证用户不掉线
	 */

	client := Client{make(chan string),adderId, adderId}
	OnlineMap[adderId] = client
	message <- getMsg(client, " Login ")

	//新开一个协程，给当前用户发送消息
	go WriteMsgToClient(client, conn)

	//新开一个协程，接收当前用户发送的信息
	go GetMsgClient(client, conn)

	for {
		select {
		case <-isQuit:
			delete(OnlineMap, adderId)
			message <- getMsg(client,  " user is quit")
		case <-hasData:
		case <-time.After(30 * time.Second):
			delete(OnlineMap, adderId)
			message <- getMsg(client,  "time out is logout")
		}
	}

}

func GetMsgClient(client Client, conn net.Conn)  {

	//存储用户发送的消息
	buf := make([]byte, 2048)

	for {
		n , err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err, "msg =", n)
			isQuit <- false
			break
		}

		msg := string(buf[:n-2])
		if strings.Contains(msg,"list") {
			for _, val := range OnlineMap {
				message <- val.Name
			}
		} else if strings.Contains(msg, "rename") {
			client.Name = strings.Split(msg,"=")[1]
			OnlineMap[client.Adder] = client
			message <- "rename is ok!!!"
		} else {
			message <- getMsg(client, msg)
		}

		hasData <- true
	}
}

func WriteMsgToClient(client Client, conn net.Conn)  {
	for msg := range client.C {
		conn.Write([]byte(msg+"\n"))
	}
}

func SendMessage()  {
	for {
		msg := <- message //没有消息的时候阻塞
		for _, cli := range OnlineMap {
			cli.C <- msg
		}
	}
}

func getMsg(client Client, str string) (msg string)  {
	msg = "【"+client.Name+"】====== " + str
	return
}

func main()  {

	//监听服务
	listener, err := net.Listen("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer listener.Close()

	go SendMessage()

	//死循环监听用户连接，服务器长期监听
	for {

		//监听连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err listener.Accept =", err)
			continue
		}

		//来了一个新连接，通过协程处理
		go HanerConn(conn)

	}


}
