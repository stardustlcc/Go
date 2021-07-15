package 基础

import (
	"fmt"
	"reflect"
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

	var dog1 *Dog
	fmt.Printf("%v\n", dog1)//nil
	dog2 := dog1
	fmt.Printf("%v\n", dog2)//nil

	if (dog1 == nil) {
		fmt.Print("true\n")//true
	} else {
		fmt.Print("false\n")
	}

	var pet Pet = dog2
	fmt.Printf("%v\n", pet)//nil
	fmt.Printf("%s\n", reflect.TypeOf(pet).String())//*main.Dog
	fmt.Printf("%T\n", dog2)//*main.Dog
	fmt.Printf("%T\n", dog1)//*main.Dog

	//只要我们把一个有类型的nil赋给接口变量，那么这个变量的值就一定不会是那个真正的nil。因此，当我们使用判等符号==判断pet是否与字面量nil相等的时候，答案一定会是false。
	if (pet == nil) {
		fmt.Print("true\n")
	} else {
		fmt.Print("false\n")//false
	}

	if (dog2 == nil) {
		fmt.Print("true\n")//true
	} else {
		fmt.Print("false\n")
	}



}
