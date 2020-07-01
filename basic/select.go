package main

import (
	"fmt"
	"time"
)

/**
Select:同时管理多个通道读写「多路复用」的语法糖。如果所有通道都不能读写，它就整体阻塞，只要有一个通道可以读写，它就会继续。

（所有跟在case关键字右边的发送语句或接收语句中的通道表达式和元素表达式都会先被求值。无论它们所在的case是否有可能被选择都会这样。
如果有一个或多个IO操作()可以完成，则Go运行时系统会随机的选择一个执行，
否则的话，如果有default分支，则执行default分支语句，
如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行）

-可处理一个或多个channel的发送或接收
-同时有多个channel时,按随机顺序处理
-可用空的select()来阻塞main函数(因为它在等待channel的输入)
-可设置超时
-可以使用break语句来终止select语句的执行。
*/

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go send(ch1, 3*time.Second)
	go send(ch2, 6*time.Second)
	recv(ch1, ch2)
}

func send(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

func recv(ch1 chan int, ch2 chan int) {
	for { // 注意这里的for循环，select并不是一个循环操作
		select {
		case v := <-ch1:
			fmt.Printf("recv %d from ch1\n", v)
		case v := <-ch2:
			fmt.Printf("recv %d from ch2\n", v)
		default: // 如果其它case条件都阻塞，则走默认分支，如果没有default则整个select会堵塞
			fmt.Printf("recv default method\n")
		}
	}
}
