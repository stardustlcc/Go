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

	_, err = c.Do("Set", "name", "cici")
	if err != nil {
		fmt.Println(err)
	}

	r, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println(r)
}
