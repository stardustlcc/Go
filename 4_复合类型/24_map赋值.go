package main

import "fmt"

func main()  {

	m1 := map[int]string{1:"mike", 2:"yoyo"}

	fmt.Println("m1=", m1)

	m1[1] = "c++"
	m1[3] = "hello" //追加，map底层自动扩容，和append一样

	fmt.Println("m1=", m1)
}
