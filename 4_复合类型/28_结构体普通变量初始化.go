package main

import "fmt"

type Student struct {
	id int
	name string
	age int
	sex byte
	adder string
}

func main()  {

	//顺序初始化，每个成员都必须初始化
	var s1 Student = Student{1, "cici", 18, 'm', "sh"}
	fmt.Println("s1=", s1)

	//指定成员初始化，没有初始化的成员自动赋值为 0
	s2 := Student{id:1, adder:"sh"}
	fmt.Println("s2=", s2)

}
