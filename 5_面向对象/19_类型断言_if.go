package main

import "fmt"

type Student struct {
	name string
	id int
}
func main()  {

	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hallo world"
	i[2] = Student{"mike", 18}

	//类型查询，类型断言
	//第一个返回的是下标，第二个是下标对应的值
	for index, data := range i {
		//第一个是值，第二个是判断的真假
		if value, ok := data.(int) ; ok == true {
			fmt.Printf("x[%d] 为 int 内容为 %d\n", index, value)
		} else if value, ok := data.(string) ; ok == true {
			fmt.Printf("x[%d] 为 string 内容为 %s\n", index, value)
		} else if value, ok := data.(Student) ; ok == true {
			fmt.Printf("x[%d] 为 string 内容为 %s\n", index, value.name)
		}
	}
}
