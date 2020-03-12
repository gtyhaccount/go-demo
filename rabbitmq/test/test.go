package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("push/fcm/fcm.go"); err != nil {
		fmt.Printf("error:%s", err)
	}
}
