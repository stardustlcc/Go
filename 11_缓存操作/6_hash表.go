package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer c.Close()

	_, err = c.Do("HSet", "books", "php", "php")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	r, err := redis.String(c.Do("HGet", "books", "php"))
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println(r)

}
