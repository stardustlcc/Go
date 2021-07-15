package main

import "fmt"

func main()  {

	//[low:high:max]
	slip := []int{1,2,3,4,5,6,7,8,9,10}
	s1 := slip[2:7:10]
	fmt.Println("s1=", s1) //[)  3,4,5,6,7
	fmt.Println("len(s1)=", len(s1))//high-low 7-2 = 5
	fmt.Println("cap(s1)=", cap(s1))//max-low 10-2 =8

	s2 := slip[:]
	fmt.Println("s2=", s2)//[0:len(slip)-1]

	//获取某个元素
	fmt.Println("一个元素：", s2[0])

	//从 0 开始 取 6 个，容量也是6
	s3 := slip[:6]
	fmt.Println("s3=", s3)

	//从slip[3]个元素开始，到结尾
	s4 := slip[3:]
	fmt.Println("s4=",s4)

}
