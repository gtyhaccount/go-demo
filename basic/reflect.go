package main

import (
	"fmt"
	"reflect"
)

type MyInt int

type people struct {
	name string
	age  int
	addr string
}

func main() {
	lee := people{
		name: "lee",
		age:  12,
		addr: "hubei wuhai guangshang",
	}

	fmt.Println(Struct2Map(lee))

    reflectTypeAlias()
    reflectFloat()

	setValue()
}

func Struct2Map(input interface{}) map[string]interface{} {
	iType := reflect.TypeOf(input)
	fmt.Printf("iType: %s \n", iType)
	fmt.Printf("iType kind(): %s \n", iType.Kind())

	iValue := reflect.ValueOf(input)
	fmt.Printf("iValue: %v \n", iValue)
	//fmt.Println("iValue:%v", iValue.Elem())

	data := make(map[string]interface{})
	for i := 0; i < iType.NumField(); i++ {
		data[iType.Field(i).Name] = iValue.Field(i)
	}

	return data
}

// 反射类型别名(自定义的类型,好处是可以给这些自定义类型加方法，典型的是string()；类比于java的基本类型和包装类)
func reflectTypeAlias(){
	fmt.Println("reflect type alias end ……")

	var mi MyInt = 1

	fmt.Println(reflect.TypeOf(mi).Name()) // MyInt
	fmt.Println(reflect.TypeOf(mi).String()) // main.MyInt
	// 返回参数的底层类型：int (即go的基本类型)
	fmt.Println(reflect.TypeOf(mi).Kind())// int
	fmt.Println(reflect.ValueOf(mi).String())// <main.MyInt Value>
	fmt.Println(reflect.ValueOf(mi).Kind()) // int
	fmt.Println(reflect.ValueOf(mi).Interface())// 1

	fmt.Println("reflect type alias end ……")
}

func reflectFloat(){
	fmt.Println("reflect float start ……")
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

	fmt.Println("reflect float end   ……")
}

/*
设置被反射对象的值
 */
func setValue(){
	fmt.Println()
	fmt.Println("set reflected variable value. start ……")

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x. 第一步
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()// 如果"第一步"里面不是取"&x"而是取"x",那么这里会panic,因为v.kind()即底层类型是float而不是ptr或者interface
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)

	fmt.Println("set reflected variable value. end ……")
}