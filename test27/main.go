package main

import (
	"flag"

	_ "github.com/garyburd/redigo/redis"
	"github.com/unicornzyn/gotest27/client"
	"github.com/unicornzyn/gotest27/server"
)

func main() {
	var key int
	flag.IntVar(&key, "k", 0, "1 client 2 server")
	flag.Parse()

	if key == 1 {
		client.Run()
	} else if key == 2 {
		server.Run()
	}
	/*
		// docker run -p 6379:6379 -d redis:5.0.7-alpine
		c, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println("connect redis faild err=", err)
			return
		}
		defer c.Close()

		_, err1 := c.Do("set", "key1", 998)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		r, err2 := redis.Int(c.Do("Get", "key1"))
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println("key1=", r)
	*/
}
