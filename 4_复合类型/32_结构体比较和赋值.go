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

	 s1 := Student{1, "cici", 18, 'm', "sh"}
	 s2 := Student{1, "cici", 18, 'm', "sh"}
	 s3 := Student{2, "cici", 18, 'm', "sh"}
	 fmt.Println("s1= s2", s1==s2)
	 fmt.Println("s1= s2", s1==s3)

	 //同类型的2个结构体变量可以相互赋值
	 var tmp Student
	 tmp = s3
	 fmt.Println("tmp =",tmp)

}
