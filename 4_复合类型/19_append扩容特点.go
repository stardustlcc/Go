package main

import "fmt"

func main()  {

	/**
		如果超过原来的容量，通常以2倍容量扩容
	 */
	a := make([]int, 0,1)
	fmt.Println("len(a)=", len(a), "cap(a)=", cap(a))

	for i:=0; i<8;i++ {
		a = append(a,i)
		fmt.Println("len(a)=", len(a), "cap(a)=", cap(a))
	}
}
