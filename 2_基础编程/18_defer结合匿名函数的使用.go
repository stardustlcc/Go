package main

import "fmt"

func main()  {

	defer func() {
		fmt.Println("aaaaaaaaa")
	}()

	fmt.Println("bbbbbbbbb")


	a := 10
	b := 20

	//defer 相当于压入栈中，10， 20 已压入进去
	defer func(a, b int) {
		fmt.Println("defer a =", a, "defer b=", b)
	}(a, b)


	a = 100
	b = 200
	fmt.Println("a=", a, "b =", b)
}
