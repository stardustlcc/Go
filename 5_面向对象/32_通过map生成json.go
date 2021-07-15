package main

import (
	"encoding/json"
	"fmt"
)

func main()  {

	//创建一个map
	m := make(map[string]interface{}, 4)
	m["Company"] = "翠翠公司"
	m["Subjects"] = []string{"php","c**","java"}
	m["IsOk"] = true
	m["Price"] = 3.14

	res , err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("res", string(res))
}
