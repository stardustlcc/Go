package main

import "fmt"
func funcA() int  {
	a := funcB()
	fmt.Println("A====", a)
	return 111
}
func funcB() int  {
	b := funcC()
	fmt.Println("B====", b)
	return 222
}
func funcC() int  {
	fmt.Println("C====", "CCC")
	return 333
}

func main()  {

	fmt.Println("start ---------------")

	a := funcA()
	fmt.Println("a====", a)

	fmt.Println("end -----------------")
}
