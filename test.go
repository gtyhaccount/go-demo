package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println(rand.Intn(1000000) + 2000000)
	fmt.Println(rand.Intn(100000) << 3)
}
