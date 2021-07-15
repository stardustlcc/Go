package main

import (
	"fmt"
	"strings"
)

func main()  {

	//Contains 字符串中是否包含另外一个字符串，包含 true 不包含 false
	fmt.Println(strings.Contains("hello world", "hello"))

	//Join 组合
	s := []string{"aaa", "bbb", "ccc"}
	str := strings.Join(s,"-")
	fmt.Println("str=", str)

	//Index 查找某个字符所在的位置，存在 具体位置的 不存 -1
	fmt.Println(strings.Index("abcd hello", "hello"))
	fmt.Println(strings.Index("abcd hello", "cc"))

	//Repeat 重复多少次
	str1 := strings.Repeat("go", 3)
	fmt.Println("str1 =", str1)

	//Split 依自定的分隔符拆分
	str2 := "aaa,bbb,ccc,ddd,eee"
	splice1 := strings.Split(str2,",")
	fmt.Println("splice1 = ", splice1)

	//Trim 去掉两头的字符
	str3 := strings.Trim("     hi my is lcc   ", " ")
	fmt.Printf("#%s#\n", str3)

	//Fields 去掉空格，将值放到切片中
	str4 := strings.Fields("  HI MY IS LCC  ")
	for _, data := range str4 {
		fmt.Println("data=", data)
	}

}
