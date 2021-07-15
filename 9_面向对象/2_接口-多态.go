package main

import "fmt"

type Dog struct {
	name string
}

type Cat struct {
	name string
}

type Sayer interface {
	Say()
}

func (dog Dog) Say()  {
	fmt.Println(dog.name, "say dog")
}

func (cat *Cat) Say()  {
	fmt.Println(cat.name, "say cat")
}

/**
	接口实现了多态
 */
func main()  {

	//接口类型变量能够存储所有实现了该接口的实例
	//不论变量传递的是结构体还是结构体指针，都是可以的
	var s Sayer
	dog := Dog{"狗狗"}
	cat := &Cat{"猫咪"}
	s = dog
	s.Say()
	s = cat
	s.Say()
}
