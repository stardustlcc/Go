package 接口

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Company string `json:"company"`
	Name string `json:"name"`
	Subject []string `json:"subject"`
}

type It2 struct {
	Subject []string `json:"subject"`
}

func main()  {


	var it = It{"公司","名称",[]string{"aaa","bbb","ccc"}}
	buf, err := json.MarshalIndent(it,"","")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(buf))

	var it1 It

	errs := json.Unmarshal([]byte(buf), &it1)
	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Printf("it1=%+v\n", it1)


	var it2 It2
	err2 := json.Unmarshal([]byte(buf), &it2)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("it2=%+v\n", it2)


}
