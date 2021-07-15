package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	/**
		设置种子
		如果种子参数一样，每次运行程序参数的随机数都一样
	 */
	//rand.Seed(99)

	/**
		采用当前系统时间作为种子参数
	 */
	rand.Seed(time.Now().UnixNano())

	//产生随机数
	for i:=0;i<5;i++{
		fmt.Println("rand=", rand.Int())
	}

	fmt.Println("--------------------")

	//100以内的值作为随机数
	for i:=0;i<5;i++{
		fmt.Println("rand=", rand.Intn(100))
	}




}
