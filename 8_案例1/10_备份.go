package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C chan string
	Name string
	Adder string
}

var OnlineMap = make(map[string]Client)

var message = make(chan string)

var hasData = make(chan bool)

var hasQuit = make(chan bool)

func main()  {

	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	go SentClientMsg()

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		go HandConn(conn)
	}

}

func HandConn(conn net.Conn)  {
	defer conn.Close()

	AdderId := conn.RemoteAddr().String()
	fmt.Println("adderId=", AdderId)

	client := Client{make(chan string), AdderId, AdderId}
	OnlineMap[AdderId] = client

	message <- makeMessage(client, "login is ok")

	//给当前的用户发送消息
	go SendUserMsg(client, conn)
	//接收用户的消息
	go ListenUserMsg(client, conn)

	for {
		select {
		case <-hasData:
		case <-time.After(30*time.Second):
			delete(OnlineMap, AdderId)
			message <- makeMessage(client, "time out is logout!!!")
		case <-hasQuit:
			delete(OnlineMap, AdderId)
			message <- makeMessage(client, "is Quit!!!!")
		}
	}

}

func makeMessage(client Client, str string) (msg string) {
	msg = "【"+client.Name+"】 "+ str +"\n"
	return
}

func SentClientMsg()  {
	for {
		msg := <-message
		for _, cli := range OnlineMap {
			cli.C <- msg
		}
	}
}

func SendUserMsg(client Client,conn net.Conn)  {
	for msg := range client.C {
		conn.Write([]byte(msg))
	}
}

func ListenUserMsg(client Client, conn net.Conn)  {
	buf := make([]byte, 2018)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			hasQuit <- false
			fmt.Println( "conn.Read or err=", err)
			break
		}

		msg := string(buf[:n-2])
		if strings.Contains(msg, "rename") {
			name := strings.Split(msg,"=")[1]
			client.Name = name
			OnlineMap[client.Adder] = client
			client.C <- makeMessage(client, "rename is ok!!")
		} else if strings.Contains(msg, "list") {
			for _, cli := range OnlineMap {
				message <- "list = "+cli.Name +"\n"
			}
		} else {
			message<- makeMessage(client, msg)
		}

		hasData <- true
	}
}