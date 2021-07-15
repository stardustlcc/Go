package main

import (
	"errors"
	"fmt"
)

func main()  {

	//var err1 error = fmt.Errorf("%s", "this is a normal error")

	err1 := fmt.Errorf("%s", "this is a normal err1")
	fmt.Println(err1)

	err2 := errors.New("this is a normal err2")
	fmt.Println(err2)
}
