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

	fmt.Println(i)

	//类型查询，类型断言
	//第一个返回下标，第二个返回下标对应的值
	for index, data := range i {
		//第一个返回值，第二个是返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("i[%d] 类型为 int 内容为 %d\n", index, value)
		}

		if value, ok := data.(string); ok == true {
			fmt.Printf("i[%d] 类型为 string 内容为 %s\n", index, value)
		}

		if value, ok := data.(Student); ok == true {
			fmt.Printf("i[%d] 类型为 Student 内容为 %v\n", index, value)
		}
	}

}
