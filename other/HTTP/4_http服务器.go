package main

import (
	"fmt"
	"net/http"
)

func HandleConn(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("req.Method=", req.Method)
	fmt.Println("req.Url=", req.URL)
	fmt.Println("req.Header=", req.Header)
	fmt.Println("req.Body=", req.Body)

	w.Write([]byte("hello go"))
}

func main()  {

	//注册处理函数，用户连接，自动调用处理函数
	http.HandleFunc("/", HandleConn)

	//监听绑定
	http.ListenAndServe(":8000", nil)
}
