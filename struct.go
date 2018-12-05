package main

import "fmt"

/*
struct:结构
-Go中没有class，结构就是为了实现部分class的功能
-type <Name> struct{} // 遵循可见性原则
-支持指向自身的指针类型成员
-支持匿名结构，可作为成员或定义成员变量
-匿名结构也可以用于map的值
-可以使用字面值对结构进行初始化
-允许通过指针来读写结构成员
-系统类型的成员可以直接拷贝赋值
-支持==和!=比较运算符，但不支持>和<
-支持匿名字段，本质上是定义了以某个类型名为名称的字段
-嵌入结构作为匿名字段看起来像继承，但不是继承
-可以使用匿名字段指针
 */

type person struct {
	Name string
	age int
}

type person2 struct { // 结构中嵌套匿名结构
	name string
	age int
	con struct { // con这个person的参数，其类型是一个匿名的结构
		phone,address string
	    }
}

type person3 struct {
	per person// 结构中嵌套另一个结构；注意：Go没有继承，这叫组合！
	// person   // 如果没有指定变量的名字，那么系统默认识别为person person,即参数名和类型名一致
	sex int
}

func main(){
	p := person{
		Name:"Lee", // 可以使用字面值设置属性；注意最后有一个逗号
	}
	p.age=11 // 可以使用“.”操作设置属性
	fmt.Println(p)

	A(p) // * 通过1和2处值的比较，知道这里传递进去的是一个值拷贝 *
	fmt.Println(p) // 1.{Lee 11}

	Aoo(&p) // 传入地址值,操作其中的对象属性
	fmt.Println(p)

	p2 := &person{ // 注意：这里直接取结构person的地址赋值给p2，p2是一个指针
		Name:"byron",
		age:23,
	}
	p2.age = 24 // 注意：这里不必使用*号取p2的地址，然后通过地址给其中对象赋值
	            //      * Go中可直接用p2.age这样操作结构中的参数 *

	// 匿名结构
	s := struct{ // 声明结构
		name string
		age int
	}{// 结构初始化
		name:"joe",
		age:16,
	}
	fmt.Println(s)

	// 结构中匿名结构的初始化
	p3 := person2{name:"joe",age:23}
	// con这个变量的类型是一个匿名的结构，所以不能在上面的大括号中直接初始化,只能用下方的这种结构初始化
	p3.con.address = "beijing"
	p3.con.phone = "12345678901"
	fmt.Println(p3)

	// 结构嵌套的初始化
	//p4 := person3{sex:1,per:person{Name:"joe123",age:13}}
	p4 := person3{}
	p4.sex = 1
	p4.per.Name = "joe123"
	p4.per.age = 13
	fmt.Println(p4)
}

func A(p person){
	p.age = 16
	fmt.Println("A()",p) // 2.A() {Lee 16}
}

func Aoo(p *person){ // 传递一个指针
	p.age = 16
	fmt.Println("Aoo() ",p)

}

