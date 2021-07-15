package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string
	Subject []string
	IsOk bool
	Price float64
}

func main()  {

	jsonStr := `
		{"Company":"123", "Subject":["php","c++"], "IsOk":true, "Price":3.14}
	`

	var tmp IT

	err := json.Unmarshal([]byte(jsonStr), &tmp)//第一个参数是结构体变量，第二个参数地址传递
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println("tmp", tmp)
	fmt.Printf("tmp= %+v\n", tmp)
}
