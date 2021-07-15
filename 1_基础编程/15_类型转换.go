package main

import "fmt"

func main()  {

	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	//布尔类型不能转换为 整型 ！！
	//fmt.Printf("flag = %d\n", int(flag))

	//0就是假，非0就是真
	//flag = 1 //不可操作 err
	//flag = bool(1) //不可操作 err

	//字符型本质上就是整型
	var ch byte
	ch = 'a'

	var t int
	//t = ch //不可操作 err
	t = int(ch) //ok success
	fmt.Println("t = ", t)

}
