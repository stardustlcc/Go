package main

import "fmt"

func main()  {

	n := [5]int{1,-10,10,3,9}
	len := len(n)
	for i:=0;i<len-1;i++{
		for j:=0; j<len-1-i;j++{
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}

	fmt.Println("n=",n)
}
