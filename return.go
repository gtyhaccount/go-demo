package main

import "fmt"

func main() {
	testReturn1()
	fmt.Println(testReturn2()) // 0
	fmt.Println(testReturn3()) // 6
	fmt.Println(testReturn4()) // 1
}

/**
退出执行，没有返回值
*/
func testReturn1() {
	fmt.Println("test return 1")
	return
}

/**
不指定返回的变量名（这里面是"a"），而且函数体中没有对返回变量名的操作，则返回值是所定义的返回变量的零值
*/
func testReturn2() (a int) {
	b := 1
	fmt.Println("testReturn2:", b)
	return
}

/**
不指定返回的变量名（这里面是"a"），函数体中有对返回变量名的操作，返回操作后的变量值
*/
func testReturn3() (a int) {
	b := 1
	a = 6 // 注意这里不是新建一个变量，而是直接拿返回变量的名称
	fmt.Println("testReturn3:", b)
	return
}

/**
指定返回的变量名（这里面是"b"），返回指定变量的值
*/
func testReturn4() (a int) {
	b := 1
	a = 6 // 注意这里不是新建一个变量，而是直接拿返回变量的名称
	fmt.Println("testReturn4:", b)
	return b
}
