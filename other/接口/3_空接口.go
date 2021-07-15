package 接口

import "fmt"

type Kong interface {

}

func print( i interface{})  {
	fmt.Println(i)
}

func main()  {
	//空接口实际上就是一个万能类型，它可用保存任意类型的值

	var i interface{} = 1
	fmt.Println(i)

	i = "你好"
	fmt.Println(i)

	var k Kong
	k = "hello"
	fmt.Println(k)

	print("你好")
}
