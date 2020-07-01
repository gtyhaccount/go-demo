package main

import "fmt"

/**
 切片：Slice （GO的动态数组的解决方案：使用 append 方法时我们需要一个源 slice 和需要附加到它里面的值。当 append 方法返回时，它返回一个新的 slice，
                                     append 方法总是增长 slice 的长度，另一方面，如果源 slice 的容量足够，那么底层数组不会发生改变，
                                     否则会重新分配内存空间。重新分配内存空间，就形成了新数组。就是所谓的“动态数组”了！）
 -其本身并不是数组，而只是指向底层数组
 -它是一种变长数组的替代方案，可以关联底层数组的局部或全部
 -它是一种引用类型（而数组是一种值类型）
 -它可以直接创建或从底层数组获取生成
 -使用len()获取元素个数，使用cap()获取容量
 -一般使用make()创建
 -如果多个slice指向相同底层数组，其中一个的值改变会影响全部

 -make([]T,len,cap)
  cap可以省略，默认和len值相同

  Slice创建方式:
	make([]type, length, capacity)
	make([]type, length)
	[]type{}
	[]type{value1, value2,..., valueN}

  Array创建方式:
	[length]Type
	[N]Type{value1, value2, ..., valueN}
	[...]Type{value1, value2, ..., valueN}

  TODO:数组和切片的区别就是是否指定了大小,以为数组是固定长度的数据类型。
*/

func main() {
	// 1.声明一个切片
	var s1 []int // 数组声明中，长度是需要显示指定的，比如：3、三个点（那么后面就需要使用大括号指定字面值了）
	fmt.Println(s1)

	// 2.声明一个切片截取数组的一部分
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 声明数组时，大括号不能忘
	fmt.Println(arr)
	s2 := arr[5:10] // 创建一个切片s2，其指向数组arr下标为5到下标为9的元素
	// s2 := arr[9] // 取数组中的一个元素
	// s2 := arr[5:] // 取数组中下标为5及之后的所有元素
	// s2 := arr[:5] // 取数组中下标为5之前的所有元素（包前不包后）
	fmt.Println(s2)

	// 3.使用make方法声明一个切片
	s3 := make([]int, 3, 10) // 声明一个指向一个int类型数组、长度为3、容量为9的切片
	// 默认长度和容量相同
	// 如果切片的长度超过容量，那么编译器会默认扩大一倍容量。
	// 因为数组是内存上连续的一片区域，所以会新分配一片内存
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	// 4.理解
	arr2 := [11]int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s4 := arr2[2:5]               // s4实际上是指向arr2数组c，d,e元素内存地址，但是s4的容量是从c开始到数组结尾
	fmt.Println(len(s4), cap(s4)) // 3(包含3个元素) 9(容量是9)

	// 5.Reslice： 从slice中重新取一次切片
	//             Reslice时索引以被slice的切片为准
	//           *** Reslice的索引不能超过被slice切片的容量,否则会越界；
	//           *** 越界不会引发底层数组内存地址的重新分配，而是会报错

	/*
		6.Append:
		  -可以在slice尾部追加元素
		  -可以将一个slice追加到另一个slice尾部
		  -如果最终长度未超过追加到的slice的容量则返回原始slice
		  -如果最终长度超过了追加到的slice，则重新分配数组，指定新的内存地址，并拷贝原始数据
	*/
	s5 := make([]int, 3, 6)
	fmt.Printf("%p\n", s5)        // "%p"-占位符，表示打印内存地址
	s5 = append(s5, 1, 2, 3)      // s5此时长度为6未超过容量；未重新分配数组内存地址
	fmt.Printf("%v %p\n", s5, s5) // %v-占位符，表示打印值
	s5 = append(s5, 1, 2, 3)      // 超过容量，重新分配地址
	fmt.Printf("%v %p \n", s5, s5)

	/*
	   举例1：一个数组两个切片有相同位置的元素，其中一个切片改变值，另外一个也会改变
	*/
	arr3 := [...]int{1, 2, 3, 4, 5}
	s6 := arr3[2:5]
	s7 := arr3[0:3]
	fmt.Println(s6, s7)
	s6[0] = 9
	fmt.Println(s6, s7)

	/*
	 7.copy：将一个slice拷贝到另一个slice中，以长度少的为准
	*/
	s8 := []int{1, 2, 3, 4, 5, 6}
	s9 := []int{7, 8, 9}
	copy(s8, s9)    // 将s9的值拷贝到s8中
	fmt.Println(s8) // [7,8,9,1,2,3]
	//copy(s9,s8) // 将s9的值拷贝到s8中
	//fmt.Println(s9)// [1,2,3]

}
