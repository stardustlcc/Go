package main

import (
	"fmt"
	"regexp"
)

func main()  {

	/**
		1) 解释规则
		2）根据规则提取关键信息
	 */

	buf := "abc azc a7c aac 888 a9c tac"
	//reg1 := regexp.MustCompile(`a.c`)
	//reg1 := regexp.MustCompile(`a[0-9]c`)
	reg1 := regexp.MustCompile(`a\dc`)
	if reg1 == nil {
		fmt.Println("reg1 = ", reg1)
	}

	result := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result=", result)

}
