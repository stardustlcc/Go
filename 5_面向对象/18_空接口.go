package main

import "fmt"

func main()  {

	//空接口，万能类型，保存任意类型的值
	var i interface{} = 1
	fmt.Println("i=", i)

	i = "abc"
	fmt.Println("i=", i)

	print(i)

}

func print(agr ...interface{})  {
	fmt.Println(agr)
}
