package main

import "fmt"
/*
其实理解闭包的最方便的方法就是将闭包函数看成一个类，一个闭包函数调用就是实例化一个类。
然后就可以根据类的角度看出哪些是“全局变量”，哪些是“局部变量”了。
*/
func main() {
	// add是一个函数类型的变量
	add:=func(base int) func(int)(int){ // 匿名函数入参是int型变量base;返回值是一个入参和出参都是int的匿名函数
		return func(i int)(int){
			return base+i
		}
	}

	add5:=add(5)
	fmt.Println("add5(10)=",add5(10))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
