package main

import "fmt"

func main()  {
	m1 := map[int]string{1:"mike", 2:"yoyo", 3:"hello"}
	delete(m1, 1)//删除key值为 1 的内容
	delete(m1, 3)//删除key值为 3 的内容
	fmt.Println(m1)
}
