package main

import "fmt"

/**
函数function:
-Go函数不支持重载、嵌套、默认参数
-特性：不定长变参（必须在参数列表的最后；func temp(a ...int)---传入值可以是多个int类型的参数）、多返回值、命名返回值参数、匿名函数、闭包
-定义函数使用关键字func，且左大括号不能另起一行
-函数也可以作为一种类型使用
*/
func main() {
	a := 1
	// temp(a)
	temp(&a) // 取内存地址作为参数传递进去
	fmt.Println(a)

	/*
	 defer:
	 -在函数体执行完成后按照多个defer声明的逆序逐个执行
	 -即使函数发生 严重错误 ，也会执行
	 -支持匿名函数的调用
	 -常用于资源清理、文件关闭、解锁、记录时间等操作

	 -Go没有异常机制，但有panic/recover模式来处理错误
	  -Panic可以在任何地方发生，但recover只有在defer调用的函数中有效
	   （因为defer可以在程序发生错误后继续执行，如果recover不声明在defer中，是没有意义的）
	*/
	// 2.defer的普通调用
	//for i:=0;i<3;i++{
	//	defer fmt.Println(i) // 它会在程序体执行完后按defer声明的逆序执行
	// 所以下面的打印语句会先执行
	//}
	//fmt.Println("程序结束")

	// 3.defer的闭包调用
	for i := 0; i < 3; i++ {
		//defer func(m int){
		//	fmt.Println(m) // 打印2，1，0；这里获取到的都是函数本身传入的参数i的值拷贝
		//}(i) // 注意：这个地方的函数，如果不加上小括号表示执行这个函数，会报错
		defer func() {
			fmt.Println(i) // 打印2，2，2;因为这里获取到的都是外面i的引用
		}() // 注意：这个地方的函数，如果不加上小括号表示执行这个函数，会报错
	}

	// 4.捕获panic
	Aoo2()

}

//func temp(a int){
func temp(a *int) { // 接受一个指针类型的参数,就可以直接操作内存上的值了
	*a = 2
	fmt.Println(a)
}

func Aoo2() {
	defer func() { // defer 的声明必须在panic之前，要不然panic都发生了，defer函数还没有被注册
		if err := recover(); err != nil {
			fmt.Println("捕获一个panic")
		}
	}()
	panic("this is a panic")
}
