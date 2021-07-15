package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int)  {
	rand.Seed(time.Now().UnixNano())
	num := 0
	for {
		num = rand.Intn(10000)
		if num > 1000 {
			break
		}
	}
	*p = num
}

func GetNum(slicp []int, num int)  {
	slicp[0] = num/1000
	slicp[1] = num%1000/100
	slicp[2] = num%100/10
	slicp[3] = num%10
}

func OnGame(slicp []int)  {

	var num int
	keySlicp := make([]int,4)

	for {

		for {
			fmt.Printf("请输入一个4位数：")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			} else {
				fmt.Println("输入的数不符合要求")
			}
		}


		GetNum(keySlicp,num)
		n := 0

		for i:=0; i<4 ;i++  {
			if keySlicp[i] > slicp[i] {
				fmt.Printf("第%d位大了一点\n", i+1)
			} else if keySlicp[i] < slicp[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		if n == 4 {
			fmt.Println("全部猜对了")
			break
		}
	}

}

func main()  {

	//产生一个随机四位数
	var randNum int
	CreateNum(&randNum)

	randSlicp := make([]int,4)
	GetNum(randSlicp, randNum)
	//fmt.Println("randNum=", randNum)
	//fmt.Println("randSlicp=", randSlicp)

	OnGame(randSlicp)

}
