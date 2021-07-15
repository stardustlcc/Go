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

	_, err2 := c.Do("expire", "name", 10)
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}



}
