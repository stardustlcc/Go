package 基础

import (
	"fmt"
	"unsafe"
)

func main()  {

	// 注意：单引号 说明这是一个字符，双引号标识一个字符串
	// byte 和 uint8 没有区别 rune 和 uint32 没有区别
	// 多出一个 byte 和 rune 只是直观上可以标识一个字符，而不仅仅是一个数字

	// byte 占 1 个字节 byte 和 uint8 没有区别
	var a byte = 65
	var b byte = 'B'
	fmt.Printf("a 的值：%c \n", a)
	fmt.Printf("b 的值：%c \n", b)


	// rune 占 4 个字节 rune 和 uint32 没有区别
	var c rune = 65
	fmt.Printf("c 的值：%c \n", c)
	fmt.Println("--------------")
	fmt.Printf("a 占了 %d 个字节 \n", unsafe.Sizeof(a))
	fmt.Printf("c 占了 %d 个字节 \n", unsafe.Sizeof(c))

	//var name byte = '你'
	//fmt.Println(name) // constant 20320 overflows byte

	var name1 rune = '你'
	fmt.Printf("name1 = %d 输出的是: %c \n", name1, name1)

	//var name2 rune = "你好啊"
	//fmt.Printf("name2 = %d 输出的是: %c \n", name2, name2)
	// cannot use "你好啊" (type untyped string) as type rune in assignment

	// 字符串
	var name2 string = "这是我的名字"
	fmt.Println(name2)

	str1 := "\r\n"
	fmt.Println(str1)
	str2 := "\\r\\n"
	fmt.Println(str2) //\r\n
	str3 := `\r\n`
	fmt.Println(str3) //\r\n

	content := `
		这个世界难不倒我
		我还是那个少年
		没有一丝丝改变
     `
	fmt.Println(content)


}