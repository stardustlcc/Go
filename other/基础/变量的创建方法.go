package 基础

import "fmt"

func main()  {
	// 定义单变量
	var str1 string = "我是中国人"
	fmt.Println(str1)

	// 定义多个变量
	var str2, str3 string
	str2 = "字符串22"
	str3 = "字符串33"
	fmt.Println(str2)
	fmt.Println(str3)

	// 声明并初始化一个变量
	str4 := "我是声明和初始化的变量"
	fmt.Println(str4)

	// 声明并初始化多个变量
	name, age := "坚果", 18
	fmt.Println(name)
	fmt.Println(age)

	// 指针变量
	name1 := "坚果"
	var pro = &name1
	fmt.Println("普通变量的值：", name1)
	fmt.Println("指针的地址：",pro)

	// 取指针的值
	fmt.Println("pro取指针的值:", *pro)

	// 创建一个匿名指针变量
	pint := new(int)
	fmt.Println("匿名指针变量的地址：", pint)
	fmt.Println("获取匿名指针变量值：", *pint)
}
