package 接口

import "fmt"

type Student struct {
	id int
	name string
}

func main()  {
	i := make([]interface{}, 3)

	i[0] = 1
	i[1] = "hello"
	i[2] = Student{10001, "cici"}
	
	for index, data := range i {
		switch valType := data.(type) {
			case int:
				fmt.Printf("i[%d] 类型为 int 内容为 %d\n", index, valType)
				break
			case string:
				fmt.Printf("i[%d] 类型为 string 内容为 %s\n", index, valType)
				break
			case Student:
				fmt.Printf("i[%d] 类型为 Student 内容为 %v\n", index, valType)
				break
		}
	}
}
