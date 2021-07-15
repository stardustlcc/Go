package main

import "fmt"

func test(m1 map[int]string)  {
	delete(m1, 1)
}
func main()  {

	m1 := map[int]string{1:"mike", 2:"yoyo", 3:"hello"}
	//map属于引用传递，底层使用的是同一个map,在函数里面删除，函数外面也被删除了
	test(m1)
	fmt.Println("m1=", m1)

}
