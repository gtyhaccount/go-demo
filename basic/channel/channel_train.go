package main

import "fmt"

type message struct {
	ID  int
	Msg string
}

func main() {
	messageChannel := make(chan *message, 256)

	go func() {
		i := 1
		for {
			i++
			messageChannel <- &message{
				ID:  i,
				Msg: fmt.Sprintf("message number:%d", i),
			}
		}
	}()

	for {
		job := <-messageChannel
		go func(m *message) {
			fmt.Printf("message:%v", m)
			fmt.Println()
		}(job)
	}

	fmt.Println("channel has over")
}
