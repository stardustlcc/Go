package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person)SetInfoValue(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p Person=", p)
}

func (p *Person)SetInfoPointer(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p *Person=", p)
}

func main()  {

	p1 := Person{"CC", 'M', 18}
	fmt.Println("p1=", p1)
	p1.SetInfoValue("dd", 'd', 11)
	fmt.Println("p1=", p1)
	(&p1).SetInfoPointer("ee",'e', 10 )
	fmt.Println("p1=", p1)

}
