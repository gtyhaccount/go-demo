package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"id":   "id",
		"age":  23,
		"name": "lee",
	}).Debug("123456")

	fmt.Println(fmt.Printf("xinge android push message error code:%v ;error content: %v", 123, "未知错误"))

}
