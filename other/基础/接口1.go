package 基础

import (
	"fmt"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetNmae(name string)  {
	dog.name = name
}

func (dog Dog) Name() string  {
	return dog.name
}

func (dog Dog) Category() string  {
	return "dog"
}


func main()  {
	dog := Dog{"pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("%v\n", ok)
	fmt.Printf("%s\n", dog.Name())//pig

	dog.SetNmae("cc")
	fmt.Printf("%s\n", dog.Name())//cc


	var pet Pet = dog
	fmt.Printf("%s\n", pet.Name())//cc
	dog.SetNmae("gg")
	fmt.Printf("%s\n", dog.Name())//gg
	fmt.Printf("%s\n", pet.Name())//cc


	dog1 := Dog{"MIMI"}
	fmt.Printf("%s\n", dog1.Name())//MIMI

	dog2 := dog1
	fmt.Printf("%s\n", dog2.Name())//MIMI
	dog1.SetNmae("DD")
	fmt.Printf("%s %s\n", dog1.Name(), dog2.Name())//DD MIMI
	dog1.name = "pp"
	fmt.Printf("%s %s\n", dog1.Name(), dog2.Name())//pp MIMI


	dog = Dog{"little pig"}
	pet = &dog
	dog.SetNmae("pp")
	fmt.Printf("%s %s\n", dog.Name(), pet.Name())//pp pp




}
