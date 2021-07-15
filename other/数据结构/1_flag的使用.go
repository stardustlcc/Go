package main

import (
	"flag"
	"fmt"
)

func main()  {

	//定义标识变量，必须通过指针访问其值
	var n   = flag.Bool("n", true, "this is a bool")
	var sep = flag.String("s", "字符串", "this is a string")

	fmt.Println(*n)
	fmt.Println(*sep)

	//更新标识变量的值，必须使用 flag.Parse() 来更新
	flag.Parse()

	*n   = false
	*sep = "我是小可爱"

	fmt.Println(*n)
	fmt.Println(*sep)
}
