package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//runtime.GOMAXPROCS(1) // 设置使用的CPU核心数
	wg.Add(2)

	go func() {
		for i := 1; i < 10000; i++ {
			fmt.Printf(" A %s \n", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i < 10000; i++ {
			fmt.Printf("B %s \n", i)
		}
		wg.Done()
	}()

	wg.Wait()
}
