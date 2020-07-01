package main

import (
	"fmt"
)

/**
map:
-类似于其它语言中的哈希表或者字典，以key-value形式存储数据
-key必须是支持==和!=运算的类型，不可以是函数、map后slice
-map使用make()创建，支持:=这种简写形式

-make([keyType]valueType,cap)

-超出容量自动扩容
-使用delete()删除键值对
-使用for range对map和slice进行迭代操作
*/
func main() {
	// 1.初始化
	var m map[int]string
	//m = map[int]string{1:"a",2:"b",3:"c"}
	m = make(map[int]string)
	m[1] = "a"
	//delete(m,1)// 目标map；要删除的map的key

	// 2.复杂map
	var m2 map[int]map[int]string
	m2 = make(map[int]map[int]string)
	m2[1] = make(map[int]string) // 如果不给m2中key为1指向的map初始化，那么就是个nil，那么下面的赋值语句会报错
	m2[1][1] = "ok"
	bc := m2[1][1]
	fmt.Printf(bc)

	// 3.利用多返回值验证键值对是否存在
	var m3 = make(map[int]string)
	fmt.Println(m3[1]) // 返回一个空的字符串

	v, ok := m3[2] // 如果只有一个返回值，那么返回value；
	// 如果有两个值，第二个值是boll类型，表示该键值对是否存在
	fmt.Println(v)
	if !ok {
		m3[2] = "123456"
	}
	v = m3[2]
	fmt.Println(v, ok)

	// 4.map的遍历
	//for i,v:=range slice{} ---如果遍历的是slice,那么返回的第一个值是index；第二个值是value
	//for k,v:=range map{} ---如果遍历的是map,那么返回的第一个值是key；第二个值是value
	arr := make([]map[int]string, 1)
	arr[0] = make(map[int]string)
	arr[0][0] = "123"
	fmt.Println(arr[0][0])
	for _, v := range arr {
		v[0] = "789"
		fmt.Println(v[0])
	}
	fmt.Println(arr[0][0])

	// 5.map的间接排序

}
