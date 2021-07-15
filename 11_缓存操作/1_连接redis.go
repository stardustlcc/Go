package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main()  {
	c, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println("err=", err)
	}

	fmt.Println(c)
	defer c.Close()
}
