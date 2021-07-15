package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool //创建redis连接池

func init()  {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:16, //最初的链接数量
		MaxActive:0,//连接池最大链接数量，不确定可以用0，0表示自定义，按需分配
		IdleTimeout:300,//链接关闭时间300秒，（300秒不使用自动关闭）
		Dial: func()(redis.Conn, error) { //要链接的数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main()  {

	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束，将连接放回连接池

	_, err := c.Do("Set", "name", "yoyo")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	r, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println(r)
	pool.Close()

}
