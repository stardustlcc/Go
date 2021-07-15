package main

import "fmt"

func main()  {
	var s interface{}
	s = 1000
	v , ok := s.(string)
	fmt.Println(ok)
	fmt.Println(v)

	v1 , ok1 := s.(bool)
	fmt.Println(ok1)
	fmt.Println(v1)

	v2 , ok2 := s.(int)
	fmt.Println(ok2)
	fmt.Println(v2)


	map1 := make(map[int]interface{})
	map1[0] = 100
	map1[1] = "你好"
	map1[2] = false
	map1[3] = 3.14

	for _, v := range map1 {
		if data, ok := v.(string); ok == true {
			fmt.Println(data, "string")
		}
		if data, ok := v.(int); ok == true {
			fmt.Println(data, "int")
		}
		if data, ok := v.(bool); ok == true {
			fmt.Println(data, "bool")
		}
		if data, ok := v.(float32); ok == true {
			fmt.Println(data, "float32")
		}
		if data, ok := v.(float64); ok == true {
			fmt.Println(data, "float64")
		}
	}

	fmt.Println("-------------------")

	//类型查询，类型断言
	//第一个返回的是下标，第二个是下标对应的值
	for _,v := range map1 {
		switch val := v.(type) { //返回的还是值
		case string:
			fmt.Println(val,  "string")
		case int:
			fmt.Println(val, "int")
		case bool:
			fmt.Println(val, "bool")
		case float64:
			fmt.Println(val, "float64")
		}
	}

	//etcd
}
