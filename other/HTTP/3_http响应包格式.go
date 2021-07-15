package main

import (
	"fmt"
	"net"
)

func main()  {
	//主动连接服务器
	conn, err:= net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	requestHeader := `
GET /go HTTP/1.1
\r\nHost: 127.0.0.1:8000
\r\nConnection: keep-alive
\r\nCache-Control: max-age=0
\r\nsec-ch-ua: " Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"
\r\nsec-ch-ua-mobile: ?0
\r\nUpgrade-Insecure-Requests: 1
\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36
\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
\r\nSec-Fetch-Site: none
\r\nSec-Fetch-Mode: navigate
\r\nSec-Fetch-User: ?1
\r\nSec-Fetch-Dest: document
\r\nAccept-Encoding: gzip, deflate, br
\r\nAccept-Language: zh-CN,zh;q=0.9
`
	//先发请求包，服务器才会回响应包
	_, writeErr := conn.Write([]byte(requestHeader))
	if writeErr != nil {
		fmt.Println("writeErr =", writeErr)
	}

	buf := make([]byte, 1024*4)
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	fmt.Printf("#%v\n#",string(buf[:n]))
}
