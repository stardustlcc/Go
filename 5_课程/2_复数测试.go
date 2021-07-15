package main

import (
	"fmt"
	"math/cmplx"
)

func main()  {
	c := 3+4i
	fmt.Println(cmplx.Abs(c))

	var a int = 3
	var b float64 = 3.14
	var d int
	d = int(float64(a) * b)
	fmt.Println(d)
}
