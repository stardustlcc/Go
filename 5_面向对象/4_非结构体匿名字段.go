package main

import "fmt"

type mystr string

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person //结构体匿名字段 ==> 继承
	int   //基础类型的匿名字段
	mystr //基础类型的匿名字段
}

func main()  {
	s1 := Student{Person{name:"ciic"},18, "hello" }
	fmt.Println("s1=", s1)

	s1.name = "yoyo"
	s1.age = 19
	s1.sex = 'm'
	s1.int = 100
	s1.mystr = "this is a string"
	fmt.Println("s1=", s1)

	s1.Person = Person{"go", 'm', 19}
	fmt.Println("s1=", s1)

}
