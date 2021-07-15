package main

import "fmt"

//定义接口类型
type Humaner interface {
	//方法，只有声明，没有实现。由别的类型（自定义类型）实现
	sayHai()
}

type Student struct {
	name string
	id int
}

//student实现了此方法
func (tmp *Student) sayHai()  {
	fmt.Println("Student tmp=", tmp)
}

type Teacher struct {
	addr string
	group string
}

//student实现了此方法
func (tmp *Teacher) sayHai()  {
	fmt.Println("Teacher tmp=", tmp)
}

type MyStr string

func (tmp MyStr) sayHai()  {
	fmt.Println("Mystr = ", tmp)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同的表现，多态
func WhoSayHi(i Humaner)  {
	i.sayHai()
}

func main()  {
	s := &Student{"student", 18}
	WhoSayHi(s)
	t := &Teacher{"TEACHER", "2222"}
	WhoSayHi(t)
	var str MyStr = "嗨，美女"
	WhoSayHi(str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = str

	//第一个参数是下标，第二个参数是下标对应的值
	for _, i := range x {
		i.sayHai()
	}

}

func main01()  {

	//定义接口类型的变量
	var i Humaner
	//只是实现了此接口的方法的类型，那么这个类型的变量（接受者类型）就可给i赋值

	s := &Student{"mike", 66}
	i = s
	i.sayHai()


	t := &Teacher{"TEACHER", "1111"}
	i = t
	i.sayHai()

	var str MyStr = "你好"
	i = str
	i.sayHai()

}