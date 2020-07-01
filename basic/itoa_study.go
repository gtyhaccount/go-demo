package main

import "fmt"

/*
iota:只能用在const的申明中，是一个从0开始的行数索引器.
*/

const (
	i = iota
	j = iota
	k // const的简写，表明和上一行等号右边的写法一致。等同于“ k = iota ”
	l
)

const m = iota
const n = iota

func main() {
	fmt.Println(i) // 0
	fmt.Println(j) // 1
	fmt.Println(k) // 2
	fmt.Println(l) // 3

	fmt.Println(m) // 0 .
	fmt.Println(n) // 0 . 不是一，表明iota只在const的括号内有用。
}
