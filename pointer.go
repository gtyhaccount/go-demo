package main

import "fmt"

/*
指针：pointer
-类似于常量和变量，指针使用前也需要声明。Go中声明格式如下：
 var var_name *var_type // *号表示这是一个指针变量
-一个指针变量可以指向任何一个值的内存地址
-Go中&表示取内存地址，放在一个变量前表示取相应变量的内存地址
 */
func main(){
	var a int = 20 // 声明普通变量
	var p *int = &a // 声明一个指向int的指针变量

	fmt.Println("a 变量的地址是：",&a)
	fmt.Println("p 变量指向的地址:", p)
	fmt.Println("*p 去变量的值：",*p)

	var pp **int // 指向指针的指针变量
	pp = &p
	fmt.Println("指向指针的指针变量的值:",pp)
	fmt.Println("取pp指向的指针变量p的值:",*pp)
	fmt.Println("取pp指向的指针变量p指向的内存地址中的值:",**pp)

	/*
	Go 语言允许函数传递指针，只需要讲函数参数设置为指针类型
	 */
	x := 100
	y := 200
	swap(&x,&y) // 取地址值传递给指针变量
	fmt.Println("x:",x,"y:",y)
}

func swap(x *int,y *int){// 传递指针类型的参数
	temp := *x
	*x = *y
	*y = temp
}
