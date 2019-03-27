package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

func main() {
	//c := make(chan bool)
	//
	//go func() {
	//	fmt.Println("goroutine 1")
	//	<-c
	//	fmt.Println("goroutine 2")
	//}()
	//
	//c <- false
	//fmt.Println("main")
	//<-c

	//str := "abcabcabc"
	//fmt.Println(strings.Replace(str, "a", "1", 3))

	timestamp := time.Now().UnixNano() / 1e6
	appKey := "9zDsNzYZk8gso8Kg0G08o4Gcw"
	sign := sha256.Sum256([]byte(appKey + strconv.FormatInt(timestamp, 10) + "ccfBb0cB831A1EC5D9e175C49a6Cf41D"))

	fmt.Println(fmt.Sprintf("%x", sign))

	//u1 := uuid.Must(uuid.NewV4())
	fmt.Println(fmt.Sprintf("%s", uuid.Must(uuid.NewV4())))
}
