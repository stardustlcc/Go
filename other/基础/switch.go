package 基础

import "fmt"

func main()  {

	num := []int{1,2,3,4,5,6,7,8,9}
	for i:= range num{
		if  i== 3 {
			num[i] |= i
		}
	}
	fmt.Println(num)

	//值类型
	num1 := [...]int{1,2,3}
	for k,v :=range num1 {
		if k == len(num1)-1 {
			num1[0] += v
		} else {
			num1[k+1] += v
		}
	}
	fmt.Println(num1)

	//引用类型
	num2 := []int{1,2,3}
	for k,v :=range num2 {
		if k == len(num2)-1 {
			num2[0] += v
		} else {
			num2[k+1] += v
		}
	}

	fmt.Println(num2)
}