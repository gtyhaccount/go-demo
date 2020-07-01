package main

import "fmt"

func main(){
	arr := [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i:",i)// 这里i都是字面值的赋值；打印：3,2,1,0
		defer func(){fmt.Println("defer_closure i:",i)}()// 这里i是引用的传递；打印4,4,4,4
		arr[i] = func(){
			fmt.Println("for i:",i)
		}
	}

	for _,v:=range arr{
		v()// 这里i是引用的传递；打印4,4,4,4
	}
}
