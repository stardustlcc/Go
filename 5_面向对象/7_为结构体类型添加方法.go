package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
} 

//带有接受者的函数叫方法
func (tmp Person) printInfo()  {
	fmt.Println("tmp =", tmp)
}

//通过一个函数，给成员赋值
func (p *Person) setInfo (n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p=", p)
}

func main()  {
	var tmp Person
	tmp = Person{"CICI", 'M', 18}
	tmp.printInfo()

	tmp1 := Person{"CC", 'S', 20}
	tmp1.setInfo("AA", 'V', 22)
	(&tmp1).setInfo("MM", 'M', 19)

}
