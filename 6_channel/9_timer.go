package main

import (
	"fmt"
	"time"
)



func main()  {

	/**
		创建一个定时器，设置时间为2s,2s后，往 timer 通道写内容（当前时间）
		时间到了，之后响应一次
	 */

	timer := time.NewTimer( 2 * time.Second )
	fmt.Println("当前时间=", time.Now())

	//2s后，往 timer.C写数据，有数据后，就可以读取

	t := <-timer.C
	fmt.Println("t=", t)


}
