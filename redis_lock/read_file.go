package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if _,err:=ioutil.ReadFile("push\\fcm\\fcm.go");err!=nil{
		fmt.Printf("err:%s",err)
	}else {
		fmt.Print("read success")
	}
}
