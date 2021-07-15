package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) printInfo()  {
	fmt.Println("p = ", p)
}

//继承了Person，成员和方法就都继承了
type Student struct {
	Person
	class string
	id int
}
func main()  {


	s1 := Student{Person{"CC", 18}, "1年级1班", 100001}
	s1.printInfo()

}
