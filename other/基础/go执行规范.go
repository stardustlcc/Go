package 基础

import (
	"fmt"
)

func main()  {
	num := 10
	sign := make(chan struct{}, num)//结构体类型的通道

	for i:=0; i<num; i++  {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}


	for j:=0; j<num; j++ {
		<-sign
	}

}
