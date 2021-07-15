package 接口

import (
	"encoding/json"
	"fmt"
)

func main()  {

	var m = make(map[string]interface{}, 4)
	m["company"] = "你好"
	m["subject"] = []string{"GO","PAYTHON"}
	m["Flower"] = true
	result, err := json.MarshalIndent(m,"","")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}
