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

	_, err = c.Do("Mset", "age", 100, "num", 200)
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	r, err := redis.Ints(c.Do("Mget", "age", "num"))
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}

}
