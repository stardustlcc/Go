package main

import (
	"encoding/json"
	"fmt"
)

//成员首字母必须大写，否则外界不能用
type It struct {
	Company string `json:"company"` //二次编码
	Subject []string `json:"subject"`
	IsOk bool	`json:",string"` //转为字符串在输出
	Price float64 `json:"-"` //此字段不会输出到屏幕
}

func main()  {

	//定义一个结构体，同时初始化
	it := It{"翠翠公司", []string{"php", "c++", "java"}, true, 3.14}

	//非格式化编码
	//res , err := json.Marshal(it)

	//格式化编码
	res, err := json.MarshalIndent(it,"", "  ")
	if err != nil {
		fmt.Println("err=", err)
	}

	fmt.Println("res", string(res))

}
