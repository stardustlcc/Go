package main

import "fmt"

type Dog struct {

}

type Sayer interface {
	Say()
}
type Mover interface {
	Move()
}

func (dog Dog) Say()  {
	fmt.Println("狗在叫")
}

func (dog Dog) Move()  {
	fmt.Println("狗在跑")
}
func main()  {

	var s Sayer
	var m Mover
	dog := Dog{}
	s = dog
	s.Say()
	m = dog
	m.Move()
}
