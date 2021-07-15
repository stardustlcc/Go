package main

import (
	"fmt"
	"net/http"
)

func main()  {

	resp, err:= http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http.Get err=", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status=", resp.Status)
	fmt.Println("Statuscode", resp.StatusCode)
	fmt.Println("Header", resp.Header)
	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err:= resp.Body.Read(buf)
		if err != nil {

		}
		if n == 0 {
			fmt.Println("read err=", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("body tmp=", tmp)
}
