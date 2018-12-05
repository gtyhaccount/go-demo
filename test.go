package main

import "fmt"

func main() {
	c := make(chan bool)

	go func() {
		fmt.Println("goroutine 1")
		<-c
		fmt.Println("goroutine 2")
	}()

	c <- false
	fmt.Println("main")
	//<-c
}
