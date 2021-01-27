package main

import "fmt"

/*
iota:只能用在const的申明中，是一个从0开始的行数索引器，同一个const种iota每出现一次，数字自动加1。
可以被认为是一个可被编译器修改的常量，在每一次const关键字出现时被重置为0。
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
