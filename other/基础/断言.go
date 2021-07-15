package 基础

import "fmt"

// 字符串切片
var con = []string{"aa", "bb", "cc"}

func main()  {

	// map
	con := map[int]string{0:"111", 1:"222", 2:"333"}
	fmt.Print(con, "\n")

	//断言判断是否是 切片
	value, ok := interface{}(con).([]string)
	fmt.Print(value, ok, "\n")

	//断言判断是否是 map
	value1, ok1 := interface{}(con).(map[int]string)
	fmt.Print(value1, ok1, "\n")

	//断言判断是否是 map
	value2, _ := interface{}(con).(map[int]string)
	fmt.Print(value2,  "\n")

	//断言判断是否是 map
	value3 := interface{}(con).(map[int]string)
	fmt.Print(value3,  "\n")

	var age int
	age = 100
	isAge, err1 := interface{}(age).(int)
	fmt.Print(isAge, err1, "\n")

	//类型转换
	var strInt = int16(-255)
	dsInt := int8(strInt)
	asInt := int(dsInt)
	asStr := string(92)
	fmt.Print(dsInt, asInt,asStr, "\n")




}