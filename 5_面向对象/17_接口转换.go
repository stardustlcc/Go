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

	//超集 可以转换为子集，反过来不行

	var per Personer //超集
	var Hum Humaner  //子集

	per = &Student{"cici", 18}

	//per = Hum  报错

	Hum = per
	Hum.sayhi()


}
