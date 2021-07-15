package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person)SetInfoValue()  {
	fmt.Println("SetInfoValue--111")
}

func (p *Person)SetInfoPointer()  {
	fmt.Println("SetInfoPointer--222")
}

func main()  {

	//结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}
	(*p).SetInfoPointer()
	p.SetInfoPointer()

	//内部的转换，先把指针 p 转成 *p 后在调用
	(*p).SetInfoValue()
	p.SetInfoValue()

}
