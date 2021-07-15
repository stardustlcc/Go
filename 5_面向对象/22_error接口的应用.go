package main

import (
	"errors"
	"fmt"
)

func Mydiv(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		res = a/b
	}
	return
}

func main()  {

	res, err := Mydiv(1, 0)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("res = ", res)
}
