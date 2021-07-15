package 基础

import (
	"fmt"
)

func main()  {

	// 整型
	var num0 int = 12
	var num1 int = 0b1100
	var num2 int = 0o14
	var num3 int = 0xC

	fmt.Printf("10进制表示的是：%d == %d \n", num0, num0)
	fmt.Printf("2进制表示的是：%b == %d \n", num1, num1)
	fmt.Printf("8进制表示的是: %o == %d \n", num2, num2)
	fmt.Printf("16进制表示的是：%X == %d \n", num3, num3)
	// 输出
	//10进制表示的是：12 == 12
	//2进制表示的是：1100 == 12
	//8进制表示的是: 14 == 12
	//16进制表示的是：C == 12

	// 浮点数
	var myfloat float32 = 10000018
	fmt.Println("浮点数：", myfloat)
	fmt.Println("浮点数加1：", myfloat+1)

	//浮点数： 1.0000018e+07
	//浮点数加1： 1.0000019e+07

	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187
	fmt.Println("myfloat01: ", myfloat01)
	fmt.Println("myfloat02: ", myfloat02)
	fmt.Println("myfloat: ", myfloat01+5)
	// 浮点数不好作比较
	fmt.Println(myfloat02 == myfloat01+5)

	//myfloat01:  1.00000184e+08
	//myfloat02:  1.00000184e+08
	//myfloat:  1.0000019e+08
	//false

}
