package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

//接口嵌套
type Animal interface {
	Sayer
	Mover
}

type Dog struct {
	name string
}

func (dog Dog) say()  {
	fmt.Println(dog.name, "汪汪叫")
}

func (dog Dog) move()  {
	fmt.Println(dog.name, "移动了")
}

/**
	接口与接口之间可以通过嵌套创造新的接口
 */
func main()  {

	var animal Animal
	animal = Dog{"花花"}
	animal.say()
	animal.move()
}
