package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWG(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Println("routinue 1 execute……")
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Println("……routinue 2 execute")
		}

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("api execute done")
}
