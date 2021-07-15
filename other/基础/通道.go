package 基础

import "fmt"

func main()  {
	ch1 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	elem1 := <- ch1
	elem2 := <- ch1
	elem3 := <- ch1
	elem4 := <- ch1

	fmt.Printf("%v\n", ch1)
	fmt.Printf("%v\n", elem1)
	fmt.Printf("%v\n", elem2)
	fmt.Printf("%v\n", elem3)
	fmt.Printf("%v\n", elem4)

	ch := make(chan []int, 1)
	s1 := []int{1,2,3}
	ch <- s1
	s2 := <- ch
	s2[0] = 100
	fmt.Print(ch, "\n")
	fmt.Print(s1, "\n")
	fmt.Print(s2, "\n")

	ch2 := make(chan [3]int, 1)
	ss1 := [3]int{1,2,3}
	ch2 <- ss1
	ss3 := <- ch2
	ss3[0] = 100
	fmt.Print(ss1, "\n")
	fmt.Print(ss3, "\n")


}
