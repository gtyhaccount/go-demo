package main

import (
	"fmt"
	"reflect"
)

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

}

func Struct2Map(input interface{}) map[string]interface{} {
	iType := reflect.TypeOf(input)
	fmt.Println("iType:%v", iType)
	fmt.Println("iType kind():%v", iType.Kind())

	iValue := reflect.ValueOf(input)
	fmt.Println("iValue:%v", iValue)
	//fmt.Println("iValue:%v", iValue.Elem())

	data := make(map[string]interface{})
	for i := 0; i < iType.NumField(); i++ {
		data[iType.Field(i).Name] = iValue.Field(i)
	}

	return data
}
