package main

import (
	"fmt"
	"strconv"
)

func main()  {

	//转换为字符串后追加到字节数组
	slice := make([]byte, 0,1024)
	slice = strconv.AppendBool(slice, true)

	//第二个数要追加的数据，第三个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)

	slice = strconv.AppendQuote(slice, "hello go")
	fmt.Println("slice =", string(slice))

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str=", str)
	fmt.Printf("%T\n", str)

	//f 指打印方式，以小数方式，-1指小数点位数（紧缩模式）64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str=", str)
	fmt.Printf("%T\n", str)

	//字符串转其他类型
	var flag bool
	var err error

	flag, err = strconv.ParseBool("true")
	if err != nil {
		fmt.Println("err= ", err)
	} else {
		fmt.Println("flag = ", flag)
	}

	//字符串转整型
	i, _ := strconv.Atoi("100")
	fmt.Printf("%T, %d\n", i,i)




}
