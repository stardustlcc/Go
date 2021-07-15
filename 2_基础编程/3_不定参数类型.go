package main

import "fmt"

func Fun1(a int, b int)  { //固定参数
	fmt.Println(a + b)
}

func Fun2(args ...int)  { //不定参数
	fmt.Println("len(args)=", len(args))

	//传统的方式
	for i := 0; i<len(args); i++ {
		fmt.Printf("args[%d] %d\n", i, args[i])
	}

	//迭代的方式
	for k, v := range args {
		fmt.Printf("k = %d, v= %d\n", k,v)
	}
}

//固定参数一定要传，不定参数根据需要传递 (注意，不定参数一定要放在形参的最后一个位置)
func Fun3(a int, args ...int)  {
	fmt.Println("a=",a)
	fmt.Println("args=", args)
}

func main()  {

	Fun1(1,2)
	Fun2(11,22,33,44,55,66)
	Fun3(1)

}
