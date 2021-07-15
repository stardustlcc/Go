package main

import "fmt"

func main()  {

	a := []int{1,2,3,4,5,6,7,8,9}

	//新切片
	s1 := a[2:5]
	fmt.Println("s1=", s1)// [3 4 5]

	a[2]= 100
	fmt.Println("a=", a)// [1 2 100 4 5 6 7 8 9]
	fmt.Println("s1=", s1)// [100 4 5]

	//另一个新切片
	s2 := s1[:2]
	fmt.Println("s2=", s2) //100 4
	s2[1] = 999
	fmt.Println("s2=", s2) // [100 999]
	fmt.Println("s1=", s1) //[100 999 5]
	fmt.Println("a=", a)   //[1 2 100 999 5 6 7 8 9]

}
