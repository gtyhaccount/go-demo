package main

import "fmt"

func main() {
	map1 := map[int]string{1:"a",2:"b",3:"c"}

	fmt.Println(map1)

	map2 := make(map[string]int)

	for k,v:=range map1{
		map2[v]=k
	}
	fmt.Println(map2)
}
