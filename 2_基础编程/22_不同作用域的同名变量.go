package main

import "fmt"

var name string = "hello"

//不同作用域，允许定义相同的变量名，使用方式采用就近原则
func main()  {

	fmt.Println("name=", name)

	name := "cici"
	fmt.Println("name=", name)

	func(){
		name := "kk"
		fmt.Println("name=", name)
	}()

	fmt.Println("name=", name)

}
