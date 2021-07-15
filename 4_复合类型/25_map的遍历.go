package main

import "fmt"

func main()  {
	m1 := map[int]string{1:"mike", 2:"yoyo", 3:"hello"}

	//map的遍历是无序的
	for key, value := range m1 {
		fmt.Printf("%d = %s\n", key, value)
	}

	//如何判断一个key值是否存在
	value, ok := m1[1]
	if ok == true {
		fmt.Println("value = ", value)
	} else {
		fmt.Println( "m1[1] == false")
	}
}
