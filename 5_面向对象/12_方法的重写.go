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
//Student也实现了这个方法，这个方法和Person方法同名，这种方法就叫重写
func (s *Student) printInfo()  {
	fmt.Println("s=", s)
}

func main()  {

	//就近原则，先找本作用域的方法，在找继承的方法
	s1 := Student{Person{"CC", 18}, "1年级1班", 100001}
	s1.printInfo()

	//显示调用继承的方法
	s1.Person.printInfo()

}
