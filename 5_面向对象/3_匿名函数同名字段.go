
package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person //匿名字段 ==> 继承
	id int
	adder string
	name string
	age int
}

func main()  {

	//自动推导类型 + 指定成员初始化
	//默认规则，就近原则，如果在本作用域找到此成员，就操作此成员
	//如果找不到，就找继承的字段
	s2 := Student{Person{name:"cici"}, 1, "sh", "sssss", 1}
	fmt.Println("s2=", s2)
	s2.id = 8
	s2.adder = "bj"
	s2.name = "gugu"

	//显示调用
	s2.Person.name = "hello"
	s2.Person.age = 18
	fmt.Println("s2=", s2)

}