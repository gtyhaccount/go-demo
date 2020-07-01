package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Goroutine:通过通信来共享内存；而不是通过共享内存来通信(比如JAVA)。

goroutine是coroutine(协程)在go语言中的实现，是一种"用户态的线程"，由用户在程序中调度。而传统意义上的"线程"是"内核态"的，其对应于操作系统上真正的线程，由系统去调度。
可以简单的将goroutine机制理解为一种“向队列中不断的插入待执行的任务”，然后真正去执行任务的还是传统意义上的线程，goroutine只是加大了线程的使用效率，避免频繁的创建、释放。
*/

func main() {
	// 如果main中这有这一部分代码，不会打印。这一点和Java区别！！！
	//go func(){
	//	fmt.Println("goroutine helloworld")
	//}()

	fmt.Println("--------------goroutine helloworld------------------------start")

	c := make(chan bool) // 创建一个传入值为bool类型的通道channel
	go func() {
		fmt.Println("goroutine helloworld")
		c <- true // 向channel中存入一个值，用来线程间通信 （1）
	}()

	<-c // 阻塞main，等待，直到取出channel里面的值，然后执行下方的代码 （2）
	/* 如果（1）和（2）的代码对调，也可以执行。因为（2）处去存，channel中的参数没有被取出来，也是会堵塞主流程的 */

	fmt.Println("--------------goroutine helloworld------------------------end")

	fmt.Println("--------------for range channel------------------------start")

	c1 := make(chan bool)
	go func() {
		fmt.Println("for range channel")
		close(c1) /* 如果这里不关闭c1，或者关闭不成功，那下面的for range会不停的循环去取channel中的值，
		   直到goroutine一直等待，形成死锁（all goroutines are asleep - deadlock!）
		*/

	}()
	for v := range c1 {
		fmt.Printf("v=%d", v)
	}

	fmt.Println("--------------for range channel------------------------end")

	fmt.Println("--------------Asynchronous execution:异步执行------------------------start")

	runtime.GOMAXPROCS(runtime.NumCPU()) //GOMAXPROCS():设置程序使用CPU个数； NumCPU()：返回计算机最大核数

	wg := sync.WaitGroup{}
	wg.Add(10) // 创建一个任务组，最多10个待完成任务

	for i := 0; i < 10; i++ {
		go f(&wg, i)
	}

	wg.Wait() // 等待所有待完成任务完成，再执行下方代码

	fmt.Println("--------------Asynchronous execution:异步执行------------------------end")
}

func f(wg *sync.WaitGroup, b int) {
	var a int
	for i := 0; i < 10000000; i++ {
		a += 1
	}

	fmt.Println(b, a)
	wg.Done() // 完成一个任务，将待完成任务数减一
}
