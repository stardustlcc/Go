package main

import "fmt"

func main()  {

	str := "abc"

	//传统的使用方式
	for i:=0; i< len(str); i++ {
		fmt.Println(string(str[i]))
	}

	//迭代的使用方式
	for key, val := range str {
		fmt.Printf("key=%d, val=%c\n", key,val)
	}

	//匿名变量 _
	for _, val := range str {
		fmt.Printf("val = %c\n", val)
	}

	//当只有一个参数的时候，表示，第二个参数丢弃
	for i := range str {
		fmt.Printf("val = %d\n", i)
	}

}
