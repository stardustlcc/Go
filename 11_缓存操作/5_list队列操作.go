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

	_, err = c.Do("lpush", "my_list", "abc","aaa","cici",100,200)
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	for {
		r, err := redis.String(c.Do("lpop", "my_list"))
		if err != nil {
			fmt.Println("err=", err)
			break
		}
		fmt.Println(r)
	}


}
