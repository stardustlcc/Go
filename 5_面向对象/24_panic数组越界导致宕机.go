package main

import "fmt"

func TestA()  {
	fmt.Println("aaaaaaaa")
}
func TestB(x int)  {
	var a [10]int
	a[x] = 100
}
func TestC()  {
	fmt.Println("cccccccc")
}

func main()  {

	TestA()
	TestB(20)
	TestC()
}
