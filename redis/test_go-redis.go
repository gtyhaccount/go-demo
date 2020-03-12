package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	rc := redis.NewClient(&redis.Options{
		Addr: "192.168.1.244:6379",
	})

	defer rc.Close()

	value := rc.Get("user-test1").Val()
	fmt.Print(value)
}
