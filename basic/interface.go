package main

import "fmt"

/*
interface：接口
-接口是一个或多个 方法签名  的集合
-只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口，这称为Structural Typing
-接口只有方法声明，没有实现， 没有数据字段
-接口可以匿名嵌入其他接口，或嵌入到结构中
-将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
-只有当接口存储的类型和对象都是nil时，接口才等于nil
-接口调用不会做receiver的自动转换
-接口支持匿名字段方法
-接口可以实现类似OOP中的多态
-空接口可以作为任何类型数据的容器
 */

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

// 实现接口
func (p PhoneConnecter) Name() string {
	return "PhoneConnecter Name()"
}

func (p PhoneConnecter) Connect() {
	fmt.Println("PhoneConnecter Connect()",p.name)
}

func Disconnect(usb USB){
	if pc,ok := usb.(PhoneConnecter);ok{
		fmt.Println(pc.name)
	}
}

func main(){
	// 用接口接收实现类
	// 如果没有实现接口声明的方法，会报错
	var u USB = PhoneConnecter{name:"PC"}
	u.Connect()

	Disconnect(u)

}
