package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	*Person //结构体匿名字段 ==> 继承
	id  int
	adder string
}

func main()  {

	s1 := Student{&Person{name:"yoyo"}, 1, "SH"}
	fmt.Println(s1.name)

	var s2 Student
	s2.Person = new(Person)
	s2.name = "hello"
	s2.adder = "bj"
	fmt.Println( s2.name, s2.adder)

}
