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

	var s1 *Student = &Student{1, "cici", 18, 'm', "sh"}
	fmt.Println("s1=", s1)
	fmt.Println("*s1=", *s1)

	s2 := &Student{id:1, adder:"bj"}
	fmt.Println("s2=", s2)
	fmt.Println("*s2=", *s2)

}
