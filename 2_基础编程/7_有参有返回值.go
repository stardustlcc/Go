package main

import "fmt"

func MaxAndMin(a,b int)(max,min int)  {
	if a > b {
		max = a
		min = b
	} else {
		min = a
		max = b
	}
	return
}

func main()  {

	max, min := MaxAndMin(10,20)
	fmt.Println("max=", max)
	fmt.Println("min=", min)

	max1, _ := MaxAndMin(100,20)
	fmt.Println("max1=", max1)

}
