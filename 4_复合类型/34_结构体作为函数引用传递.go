package main

import "fmt"


type Student struct {
	id int
	name string
	age int
	sex byte
	adder string
}

func test02(s1 *Student)  {
	s1.id = 2
	fmt.Println("test02 =", s1)
}


func main()  {
	s1 := Student{1, "cici", 18, 'm', "sh"}
	fmt.Println("s1=", s1)
	test02(&s1)//形参无法改实参
	fmt.Println("s1=", s1)
}
