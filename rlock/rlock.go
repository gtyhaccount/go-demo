package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//全局变量
var count int
var rLock sync.RWMutex

func Read(i int) {
	rLock.RLock()
	fmt.Printf("读 goroutine%d 数据=%d\n", i, count)
	defer rLock.RUnlock()
}

func Write(i int) {
	rLock.Lock()
	count = rand.Intn(1000)
	fmt.Printf("写 goroutine%d 数据=%d\n", i, count)
	defer rLock.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go Write(i)
	}
	for i := 0; i < 5; i++ {
		go Read(i)
	}
	time.Sleep(time.Second * 2)
	//执行结果：
	/*
	   写 goroutine0 数据=81
	   写 goroutine1 数据=887
	   读 goroutine1 数据=887
	   写 goroutine3 数据=847
	   写 goroutine4 数据=59
	   读 goroutine0 数据=59
	   读 goroutine3 数据=59
	   写 goroutine2 数据=81
	   读 goroutine4 数据=81
	   读 goroutine2 数据=81
	*/
}
