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
	//指针有合法指向后，才操作成员
	var s1 Student
	//定义一个指针变量，保存s1的地址
	p1 := &s1
	p1.id = 1
	p1.name = "cici"
	(*p1).age = 18
	fmt.Println(p1)

	//2、通过new 申请一个结构体(通过new给他申请一块内存)
	p2 := new(Student)
	p2.id = 1
	p2.name = "hello"
	fmt.Println("p2=", p2)
}
