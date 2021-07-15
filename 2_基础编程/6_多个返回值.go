package main

import "fmt"

func myFun01() (int, int, int) {
	return 1,2,3
}


func myFun02() (a int, b int, c int) {
	a,b,c = 11,22,33
	return
}

func main()  {
	a, b, c := myFun01()
	fmt.Println("a=", a, "b=", b, "c=", c)

	i, j, k := myFun02()
	fmt.Println("i=", i, "j=", j, "k=", k)
}
