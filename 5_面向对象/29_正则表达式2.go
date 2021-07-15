package main

import (
	"fmt"
	"regexp"
)

func main()  {

	str := "3.14 567 abc 1.23 7. 8.99 1sfafad 7.88 0.01 10000.0001"
	//获取有效的小数   + 是匹配前一个字符的1次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("error=", reg)
		return
	}
	result := reg.FindAllStringSubmatch(str, -1)
	fmt.Println("result=", result)

}
