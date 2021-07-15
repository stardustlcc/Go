package main

import "fmt"

type User struct {
	Name string
	age int
}

/**
	user 结构体实现的方法
	值传递
 */
func (user User) Notify(name string)  {
	user.Name = name
	fmt.Println("user", user)
}

func (user *User) NotifyPoint (name string)  {
	user.Name = name
	fmt.Println("user *user=", user)
}

func Test(name string)  {
	fmt.Println("name=", name)
}

/**
	方法实现了继承
 */
func main()  {

	var user = User{"nita", 18}
	fmt.Println("user main=", user)
	//值传递，不会对之前的结构体的值影响
	user.Notify("cici")
	fmt.Println("user main=", user)


	//指针传递，可以对主函数里面的结构体进行修改
	user1 := &User{"aaa", 11}
	user1.NotifyPoint("bbbb")
	fmt.Println("use1 main=", user1)

	//普通函数的参数不能接收指针类型的参数
	 var name string = "hello"
	//Test(&name) 普通函数不能将
	Test(name)
}
