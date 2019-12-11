package main

import "fmt"
// Go 语言中数组是固定长度的数据类型，它包含相同类型的连续的元素，这些元素可以是内建类型，像数字和字符串，也可以是结构类型，元素可以通过唯一的索引值访问，从 0 开始。
// 数组是很有价值的数据结构，因为它的内存分配是连续的，内存连续意味着可是让它在 CPU 缓存中待更久，所以迭代数组和移动元素都会非常迅速。

// Go语言中不像其它语言为了节省空间，而将数组当做参数传递时是引用传递；Go中传递数组时是一种值传递，即将数组拷贝一份传递
// 要使用数组的引用传递，可以使用"切片"
func main() {
	//a := [1]int{}
	//b := [2]int{}
	//a = b // *数组的长度也是数组类型的一部分
	      //  所以不同长度的数组是不能直接比较和赋值的
	      // 这里直接赋值会报错“cannot use b (type [2]int) as type [1]int in assignment”
	//fmt.Println(a)

	var arr1 = [2]int{}
	fmt.Println(arr1)

	// 显示的使用下标索引制定该数组第20个元素为1，其他元素为零值
	var arr2 = [20]int{19:1}
	fmt.Println(arr2)

	// "..."表示Go自动去计算这个数组的长度
	var arr3 = [...]int{1:2,19:6}// 最少要20个元素
	fmt.Println(arr3)

	// 取数组arr3的内存地址给类型为指针类型的变量p，打印p和直接打印arr3效果一样
	var p *[20]int = &arr3
	fmt.Println(p)

	// 存放指针的数组("*"表示指针；“&”表示取内存地址)
	var x int = 1
	var y int = 2
	p_arr := [...]*int{&x,&y}// 存放的是两个int类型对象的内存地址值
	fmt.Println(p_arr)

	// new 关键字返回的是一个指向数组的指针
	arr4 := new([3]int)
	fmt.Println(arr4)

	// 多维数组
	arr5 := [2][3]int{{1,1,1},{2,2,2}} // 长度为2的存放元素为（长度为3的int类型数组）的数组
	arr5[1][2] = 3
	fmt.Println(arr5)

}
