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
	//定义一个普通变量
	var s1 Student
	//操作成员，需要使用（.）运算符
	s1.id = 1
	s1.name = "cici"
	s1.age = 18
	s1.sex = 'm'
	s1.adder = "sh"
	fmt.Println("s1=", s1)

}
