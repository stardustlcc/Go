package main

import (
	"encoding/json"
	"fmt"
)

func main()  {

	jsonStr := `
		{"Company":"123", "Subject":["php","c++"], "IsOk":true, "Price":3.14}
	`

	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println("err=", err)
	}

	fmt.Printf("m= %+v\n", m)

	str := m["Company"]
	fmt.Println("str=", str)

	var str1 string
	//str1 = m["Company"] 类型不一致
	//fmt.Println("str1=", str1)

	//类型断言
	for key, val := range m {
		//类型不一致，也不行
		//if key == "Company" {
			//str1 = val
			//fmt.Println("str=", str1)
		//}

		switch data := val.(type) {
		case string:
			str1 = data
			fmt.Printf("map[%s] ===> data=%s\n", key, str1)
		case bool:
			fmt.Printf("map[%s] ===> data=%t\n", key, data)
		case float64:
			fmt.Printf("map[%s] ===> data=%f\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s] ===> data=%v\n", key, data)
		}
	}

}
