package 基础

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string)  {
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
	fmt.Printf("%v\n", ok) //false

	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("%v\n", ok) //true


	var pet Pet = &dog
	fmt.Printf("%s \n %q\n", pet.Category(), pet.Name())


}
