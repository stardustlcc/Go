package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct {

}

type Person struct {

}

func (dog Dog) Move()  {
	fmt.Println("狗狗在动")
}

func (per Person) Move()  {
	fmt.Println("小偷别跑")
}

func main()  {

	var d Mover = Dog{}
	d.Move()
	var p Mover = Person{}
	p.Move()

}
