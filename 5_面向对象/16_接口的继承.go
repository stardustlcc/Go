package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type Personer interface { //超集
	Humaner   //匿名字段，相当于继承了这个方法
	sing(str string) //有自己的方法
}

type Student struct {
	name string
	id int
}

func (s *Student) sayhi()  {
	fmt.Println("student = ", s)
}

func (s *Student) sing(str string)  {
	fmt.Println(s.name+"在唱歌"+str)
}

func main()  {

	//定义一个接口类型的变量
	var i Personer

	s := &Student{"cici", 18}
	i = s
	i.sing("我是一只笑笑小小鸟")
	i.sayhi()


}
