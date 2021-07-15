package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person //匿名字段 ==> 继承
	id int
	adder string
}

func main()  {

	//顺序初始化
	var s1 Student = Student{Person{"cici", 'm', 18}, 1, "bj"}
	fmt.Println("s1=", s1)

	//自动推导类型 + 指定成员初始化
	s2 := Student{Person{name:"cici"}, 1, "sh"}
	fmt.Println("s2=", s2)
	//%+v 显示更详细
	fmt.Printf("s2=%+v\n", s2)
}
