package 接口

import "fmt"

type Student struct {
	id int
	name string
}

type Human interface {
	say(str string)
}

type Person interface {
	Human
	sing(str string)
}

func (tmp *Student) say(str string)  {
	fmt.Printf("Student say [%s] say:%s\n", tmp.name, str)
}

func (tmp *Student) sing(str string) {
	fmt.Printf("Student sing [%s] say:%s\n", tmp.name, str)
}

func main()  {

	var p Person
	s := &Student{100000001, "欢欢"}
	p = s
	p.sing("我是施欢欢")
	p.say("我是小可爱")


	var myP Person //超集
	var myH Human  //子集

	//超集可以转换为子集，子集不可用转为超级
	s1 := &Student{20001, "翠翠"}
	myP = s1
	myP.say("我是小可爱")
	myP.sing("我是可爱多")

	s2 := &Student{30001, "小弥"}
	myH = s2
	myH.say("可盐可咸的小可爱")
	//myH.sing("1111")//myH.sing undefined (type Human has no field or method sing)



}
