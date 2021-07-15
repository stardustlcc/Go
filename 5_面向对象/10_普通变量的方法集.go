package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) setInfo()  {
	fmt.Println(1111)
}

func main()  {

	p := Person{"CICI", 18}
	p.setInfo()

}
