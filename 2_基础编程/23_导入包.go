package main

//忽略此包
import _ "time"

//给包名取别名
import osTmp "os"

//调用函数，无需通过包名
import . "fmt"

//引入多包

func main()  {

	Println("hai boy")
	println(osTmp.Args)

}
