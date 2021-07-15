package main
import "fmt"

var global *int

func P ()  {
	var s int
	s = 200
	global = &s
}

func main()  {

	p := new(int)
	*p = 100
	fmt.Println(*p)

	str := new(string)
	*str = "我是小可爱"
	fmt.Println(*str)

	//变量s从局部函数中逃逸了
	P()
	fmt.Println(*global)

	var str1 = []string{"你好","hello"}
	str1[0] = "world"
	//str1[1] = true
	fmt.Println(str1)


}
