package 接口

import "fmt"

//结构体
type Student struct {
	id int
	name string
}

//接口
type Humaner interface {
	sayhi()
}

//接口
type Personer interface {
	Humaner
	sing(str string)
}

//接口的实现
func (tmp *Student) sayhi()  {
	fmt.Printf("Student[%s %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(str string)  {
	fmt.Printf("Student在唱歌：%s\n", str)
}


func main()  {
	var p Personer
	s := &Student{100001, "cici"}
	p = s
	p.sing("我是一只小小小小鸟")
	p.sayhi()


}