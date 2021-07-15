package main

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	C chan string
	Name string
	Adder string
}

var OnlineMap = make(map[string]Client)

var message = make(chan string)

var isQuit  = make(chan bool)

func main()  {

	listener, error := net.Listen("tcp", "127.0.0.1:8088")
	if error != nil {
		fmt.Println("error=", error)
		return
	}

	defer listener.Close()

	go sendClientMessage()

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			continue
		}

		go HandConn(conn)

	}
}

func HandConn(conn net.Conn)  {

	defer conn.Close()

	adderId := conn.RemoteAddr().String()

	fmt.Println("adderId=", adderId)

	client := Client{make(chan string), adderId, adderId}
	OnlineMap[adderId] = client

	message <- makeClientMessage(client, " login ")

	//给所有在线的用户发送信息
	go writeMsgClient(client, conn)

	go getClientMessage(client,conn)

	for {
		select {
		case <-isQuit:
			delete(OnlineMap, adderId)
			message<- "this user is leave!!"
		}
	}

}

func makeClientMessage(client Client, str string) (msg string)  {
	msg = "【 "+ client.Name +"】 " + str + "\n"
	return
}

func sendClientMessage()  {
	for {
		msg := <-message
		for _, cli := range OnlineMap {
			cli.C <- msg
		}
	}
}

func writeMsgClient(client Client, conn net.Conn)  {
	for msg := range client.C {
		conn.Write([]byte(msg))
	}
}

func getClientMessage(client Client, conn net.Conn)  {

	buf := make([]byte,2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			isQuit <- false
			break
		}
		msg := string(buf[:n-2])
		//rename=cici
		if strings.Contains(msg,"rename") {
			name := strings.Split(msg,"=")[1]
			client.Name = name
			OnlineMap[client.Adder] = client

		} else if strings.Contains(msg,"list") {
			for _, cli := range OnlineMap {
				message <- "list info: "+ cli.Name +"\n"
			}
		} else {
			message <- makeClientMessage(client, msg)
		}

	}
}