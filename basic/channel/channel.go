package main

import (
	"fmt"
	"time"
)

/**
Channel:（可以理解成java中的BlockingQueue）
-Channel是goroutine的沟通桥梁，大多是阻塞同步的（也可以是异步的）。
-通过make创建,第一个类型参数限定通道可以容纳的数据类型，再提供第二个整数参数作为通道的容器大小;
 大小参数是可选的，如果不填，那这个通道的容量为零，叫着「非缓冲型通道」，非缓冲型通道 必须 确保有协程正在尝试读取当前通道，否则写操作就会阻塞直到有其它协程来从通道中读东西。
-通过close关闭。
-Channel是引用类型。
-可以通过for range来迭代操作channel.
-可以设置单向或双向的通道（channel）。
-可以设置有缓存的通道(channel)

通道满了，写操作就会阻塞，协程就会进入休眠，直到有其它协程读通道挪出了空间，协程才会被唤醒。如果有多个协程的写操作都阻塞了，一个读操作只会唤醒一个协程。

通道空了，读操作就会阻塞，协程也会进入睡眠，直到有其它协程写通道装进了数据才会被唤醒。如果有多个协程的读操作阻塞了，一个写操作也只会唤醒一个协程。
*/
func main() {
	//1. 创建非缓冲通道,即代表channel中不能缓冲元素，有存放的动作就必须有  goroutine 去取,同样,有取的动作就必须有goroutine去存，必须配对。
	c1 := make(chan int)

	go func() {
		fmt.Println(<-c1) // 确保通道c1有协程去读，否则(2)处的代码会报错"fatal error: all goroutines are asleep - deadlock!"
		// 即使该goroutine的声明在(2)的下面也会报同样的错
	}()

	c1 <- 1 // (2)

	fmt.Println("1……")

	c2 := make(chan int, 3) // 创建可以缓冲3个元素的通道

	//2. 如果channel中的元素已满，再次存放，会报错"fatal error: all goroutines are asleep(睡着的) - deadlock!"
	for i := 0; i < 3; i++ {
		c2 <- i
		time.Sleep(10)
	}

	/*
		3.Go 语言的通道有点像文件，不但支持读写操作， 还支持关闭。读取一个已经关闭的通道会立即返回通道类型的「零值」，
		  而写一个已经关闭的通道会抛异常（确保通道写安全的最好方式是由负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。）。
		  如果通道里的元素是整型的，读操作是不能通过返回值来确定通道是否关闭的。
	*/
	fmt.Println("3……")
	var ch = make(chan int, 4)
	ch <- 1
	ch <- 2
	close(ch)

	value := <-ch
	fmt.Println(value)
	value = <-ch
	fmt.Println(value)
	value = <-ch
	fmt.Println(value)
	/*
		-------
		1
		2
		0
	*/

	/*
		4.for range 语法我们已经见了很多次了，它是多功能的，除了可以遍历数组、切片、字典，还可以遍历通道，取代箭头操作符。
		  当通道空了，循环会暂停并阻塞，当通道关闭时，阻塞停止，循环也跟着结束了。当循环结束时，我们就知道通道已经关闭了

		  通道如果没有显式关闭，当它不再被程序使用的时候，会自动关闭被垃圾回收掉。不过优雅的程序应该将通道看成资源，显式关闭每个不再使用的资源是一种良好的习惯。
	*/
	fmt.Println("4……")
	var ch2 = make(chan int, 4)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	// for range 遍历通道
	for value := range ch2 {
		fmt.Println(value)
	}
}
