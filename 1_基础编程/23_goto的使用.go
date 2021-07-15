package main

import "fmt"

func main()  {

	/**
		goto 可以用在任何地方，不过不能夸函数使用
	**/

	fmt.Println("111111111")
	goto End //goto是关键字，End是用户起的名字,叫标签
	fmt.Println("222222222")
	End:
	fmt.Println("333333333")

}
