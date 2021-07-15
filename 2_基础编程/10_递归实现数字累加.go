package main

import "fmt"

//实现 1+ 2 + 3+ ...100

//传统的方式
func test01() (sum int) {
	for i:=1;i<=100;i++ {
		sum +=i
	}
	return
}

//通过递归来实现
func test02(i int) (sum int) {
	if i == 1 {
		return 1
	}
	return i + test02(i - 1)
}

func main()  {

	var sum int

	sum = test01()
	fmt.Println("sum=",sum)

	sum = test02(100)
	fmt.Println("sum=",sum)

}
