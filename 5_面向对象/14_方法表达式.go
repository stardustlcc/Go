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


	//方法表达式
	f := (*Person).setInfo
	f(&p)  //显示吧接受者传递过去  p.setInfo()

	f1 := (Person).info
	f1(p)  //显示吧接受者传递过去  p.info()

}