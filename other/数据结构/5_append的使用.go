package main

import "fmt"

func main()  {

	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	fmt.Println(x)

	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Sprintf("%q\n",string(runes))


	fmt.Println([]rune("hellow,世界"))
}
