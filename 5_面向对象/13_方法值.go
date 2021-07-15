package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) setInfo()  {
	fmt.Printf("%p, %v\n", p, p)
}

func (p Person) info()  {
	fmt.Println("p=", p)
}

func main()  {
	p := Person{"cc", 18}
	fmt.Printf("%p, %v\n", &p, p)
	p.setInfo() //传统的调用方式

	pFunc := p.setInfo  //这个就是方法值
	pFunc() 			//等价于 p.setInfo()

	vFunc := p.info
	vFunc()				//等价于 p.info()

}