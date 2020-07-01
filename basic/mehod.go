package main

import "fmt"

/*
method:方法
-通过显式指定receiver（即方法声明时，方法名前面括号中指定的类型）来实现与某个类型的组合
- * 只能为同一包中的类型定义方法 *
-receiver可以是类型的值或者指针
-Go中不存在方法重载
-可以使用值或指针来调用方法，编译器会自动完成转换
- * 如果外部结构和嵌入结构存在同名方法，优先使用外部结构的方法 *
- * 类型别名不会拥有底层类型所附带的方法 *
-方法可以调用结构中的非公开方法
 */

type TZ int // 声明一个类型TZ，它的底层类型是int，但它并不拥有int类型的方法，
            // 因为int并不是这个包中的类型
            // 通过这种声明，可以给普通类型添加一些自定义的高级操作

type A struct {
	name string
}

type B struct {
	name string
}

func main(){
	a := A{name:"A"}
	a.print()
	fmt.Println(a)

	b := B{name:"B"}
	b.print()
	fmt.Println(b)

	var tz TZ
	tz.incase()
	fmt.Println(tz)
}

func (a *A) print(){ // 绑定的是指针对象，所以可以改变指针指向内存地址中对象的值
	              // 反之，B的print()不能改变其值
	a.name = "AA"
	fmt.Println("A")
}

// Go中没有方法重载的概念
// 即使方法签名不同，Go中同一个结构的方法名也不能相同
//func (a A) print(i int){
//	fmt.Println("A")
//}

func (b B) print(){
	b.name = "BB"
	fmt.Println("B")
}

func (tz TZ) print(){ // 为类型绑定方法
	fmt.Println("TZ")
}

func (tz *TZ) incase(){
	*tz += TZ(100) // 虽然TZ底层是int，但它不能和int不是同一种类型，需要强制转换
	                  // 很显然，TZ不能继承int的方法，那么很显然不是同一种类型
}
