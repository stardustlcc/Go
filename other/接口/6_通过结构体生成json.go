package 接口

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Company string `json:"company"`
	Subject []string `json:"subject"`
	Isok bool `json:"isok"`
	Price float64 `json:"price"` //二次编码
	Test string `json:"-"`  //此字段不会被输出
	Hello bool `json:",string"` //转为字符串类型
}

func main()  {
	var s = It{"Company", []string{"go","c++","Python"}, true, 34.23, "我是小可爱",false}

	//返回切片
	//buf, err := json.Marshal(s)

	//格式化编码
	buf, err := json.MarshalIndent(s,"","")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))


}
