package main

import (
	"container/list"
	"fmt"
)

func main() {

	t := new(Topic)
	t.Name = "lee"
	t.chs = list.New()
	t.chs.PushBack("t.u.123456")

	for e := t.chs.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
}

type Topic struct {
	Name string

	chs *list.List
}
