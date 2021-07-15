package main

import "fmt"

func funTest(a int) {
	if  a == 1  {
		fmt.Println("a == 1")
		return
	}
	funTest(a-1)
	fmt.Println("a ==", a)
}

func main()  {
	funTest(10)
}
