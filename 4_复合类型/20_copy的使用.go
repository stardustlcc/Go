package main

import "fmt"

func main()  {

	srcSlice := []int{1,2,3}
	desSlice := []int{9,9,9,9,9,9,9,9}

	copy(desSlice, srcSlice)
	fmt.Println("desSlice=", desSlice)//[1 2 3 9 9 9 9 9]

}
