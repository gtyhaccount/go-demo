package main

import (
	"time"
)

type message struct {
	msg string
}

var ch = make(chan message)

func main() {
	//fmt.Println("Hello, World!")
	//time.Now().Unix()
	//fmt.Println(time.Now().UnixNano()/1e6)
	//str:="1970-01-01 08:00:00"
	//date, _ := time.Parse("2006-01-02 15:04:05", str)
	//fmt.Println(date.UnixNano()/1e6)

	// 启动一个 goroutine 发送数据
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- message{
				msg: "str",
			}
			time.Sleep(time.Second) // 模拟一些工作
		}
	}()

	// 接收数据
	//for value := range ch {
	//	fmt.Println(value.msg) // 输出顺序将是 1, 2, 3, 4, 5
	//}

}
