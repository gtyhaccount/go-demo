package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Id   int64
	Age  int64
	Name string
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) getId() int64 {
	return p.Id
}

func (p *Person) Error() string {
	return "Person error"
}

func TestHW(t *testing.T) {
	author := "Lee"
	fmt.Println(reflect.TypeOf(author))
	fmt.Println(reflect.ValueOf(author))

	person := Person{
		Id:   1,
		Age:  23,
		Name: "Lee",
	}

	pType := reflect.TypeOf(person)
	fmt.Printf("type:%s \n", pType)
	fmt.Printf("type name:%s \n", pType.Name())
	fmt.Printf("type name string:%s \n", pType.String())
	// Kind本质是一个无符号的int，返回值是一个字符串，类似于枚举
	fmt.Printf("type kind:%s \n", pType.Kind()) // struct
	fmt.Printf("field of special index:%v \n", pType.Field(0))
	fmt.Printf("type method:%v \n", pType.NumMethod())

	pValue := reflect.ValueOf(person)
	fmt.Printf("value:%s \n", pValue.Kind())
	fmt.Printf("value:%s \n", pValue.String())
	fmt.Printf("value:%s \n", pValue.Field(1).String())
	// 将表示反射值的Value对象转换成go的interface对象
	fmt.Printf("value:%v \n", pValue.Interface())
	fmt.Printf("value:%v \n", pValue.Interface().(Person).Name)
}

// 修改反射对象的值
func TestUpdateReflectObjectValue(t *testing.T) {
	i := 1
	// 获取i指针的反射Value对象
	v := reflect.ValueOf(&i)
	// Elem()获取指针指向的变量
	v.Elem().SetInt(10)
	fmt.Println(i)
}

// 测试类型是否实现了某些接口
func TestIsImplementsInterface(t *testing.T) {
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&Person{})
	customError := reflect.TypeOf(Person{})

	fmt.Println(customError.Implements(typeOfError))    // false
	fmt.Println(customErrorPtr.Implements(typeOfError)) // true 说明是*Person实现了error接口
}
