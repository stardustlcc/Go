package main

import "fmt"

func main()  {

	//定义别名
	type bigint int64

	var a bigint
	fmt.Printf( "%T\n", a)

	type (
		long int64
		char byte
	)

	var c long = 10
	var ch char = 'a'

	fmt.Printf("ch = %c\n a= %d\n", ch, c)

}
