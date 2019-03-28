package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 1; i < 100; i++ {
			fmt.Println("test get goroutine id:", GetGoroutineId())
		}
		wg.Done()
	}()

	wg.Wait()
}

/**
利用runtime.Stack的堆栈信息，将当前的堆栈信息写入到一个slice中，
堆栈的第一行为 “goroutine #### […”，其中“####”就是当前的Goroutine Id.
*/
func GetGoroutineId() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover:panic info:%v \n", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
