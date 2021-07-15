package main

import "fmt"

func main()  {
	for i, r := range "hello,world，世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r,r)
	}

	println(len(","))
}
