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

	//自动推导类型 + 指定成员初始化
	s2 := Student{Person{name:"cici"}, 1, "sh"}
	fmt.Println("s2=", s2)
	s2.id = 8
	s2.adder = "bj"
	s2.Person.name = "hello"
	fmt.Println("s2=", s2)

}