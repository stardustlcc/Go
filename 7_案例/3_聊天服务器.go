package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//用户信息结构体
type Client struct {
	C chan string 	//存储用户的管道
	Name string 	//用户名
	Adder string 	//存储用户地址
}


//存储在线的用户连接,需要申请空间
var OnlineMap map[string]Client

//需要通信的一个管道
var message = make(chan string)

//对方是否主动退出
var isQuit = make(chan bool)

//对方是否有数据发送
var hasData = make(chan bool)


//给当前用户发送信息
func WriteMsgToClient(client Client, conn net.Conn)  {
	for msg := range client.C { //给当前客户端发送信息
		conn.Write([]byte(msg+"\n"))
	}
}

func GetUseSendMessage(client Client, conn net.Conn)  {

	buf := make([]byte, 2048)
	for {
		n ,err := conn.Read(buf)
		if n == 0 { //对方断开，或出问题
			fmt.Println("conn.read err=", err)
			isQuit <- false
			return
		}

		msg := string(buf[:n-2]) //mac测试多2个换行

		//过滤用户的信息
		if strings.Contains(msg,"list") {
			conn.Write([]byte("\n\nuser list:\n\n"))
			for key, onlineUser := range OnlineMap {
				msg := key +":"+onlineUser.Name + "\n"
				conn.Write([]byte(msg))
			}
		} else if strings.Contains(msg, "rename") {
			client.Name = strings.Split(msg,"=")[1]
			OnlineMap[client.Adder] = client
			conn.Write([]byte("rename ok!\n\n"))
		} else {
			//接收到的用户信息，进行转发
			message <- getMessage(client, msg)
		}

		hasData <- true
	}
}

func getMessage(client Client, msg string) (str string) {
	str = " name = "+ client.Name + "(IPAdder="+client.Adder+") " + msg
	return
}

//处理用户的连接
func HandleConn( conn net.Conn)  {
	defer conn.Close()

	//获取网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建链接用户信息结构体
	client := Client{make(chan string), cliAddr, cliAddr}
	//将在线的用户信息添加到map
	OnlineMap[cliAddr] = client

	//广播，某个用户上线了
	message <- getMessage(client, "Login")

	//新开一个协程专门给当前客户端发送信息
	go WriteMsgToClient(client, conn)



	//新开一个协程，接收用户发送过来的数据
	go GetUseSendMessage(client, conn)



	for {
		//通过select检测channel流动
		select {
		case <-isQuit://如果有用户下线了才不会阻塞
			//删掉当前用户
			delete(OnlineMap, cliAddr)
			message<-getMessage(client, "logout!")

		case <-hasData:

		case <-time.After(10*time.Second): //60秒后超时
			delete(OnlineMap, cliAddr)
			message <- getMessage(client, "time out leave out logout!")

		}
	}

}

func Message()  {
	for {
		//没有消息前，这里会阻塞
		msg := <- message
		//遍历在线用户，给每个用户发送消息
		for _,client := range OnlineMap {
			client.C <- msg
		}
	}
}

func main()  {

	//服务器监听
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	//给map分配空间
	OnlineMap = make(map[string]Client)

	//新开一个协程，转发消息，只要有消息来了，就要遍历map,给map每个成员都发送消息
	go Message()

	//主协程，循环阻塞等待用户的链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			continue
		}

		//处理用户连接
		go HandleConn(conn)
	}

}
