package main

/**
defer用于资源的释放，会在函数返回之前进行调用。一般采用如下模式：

	f,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

如果有多个defer表达式，调用顺序类似于栈，越后面的defer表达式越先被调用。

****** 函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。 ******

注意:其实使用defer时，用一个简单的转换规则改写一下，就不会迷糊了。改写规则是将return语句拆成两句写，return xxx会被改写成:

	返回值 = xxx
	调用defer函数
	空的return


*/
func main() {
	println(f1())
	println(f2())
	println(f3())
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// f1()改写
//func f1() (result int) {
//	result = 0  //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
//	func() { //defer被插入到return之前执行，也就是赋返回值和ret指令之间
//		result++
//	}()
//	return
//}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// f2()改写
//func f2() (r int) {
//	t := 5
//	r = t //赋值指令
//	func() {        //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
//		t = t + 5
//	}
//	return        //空的return指令
//}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

// f3()改写
//func f3() (r int) {
//	r = 1  //给返回值赋值
//	func(r int) {        //这里改的r是传值传进去的r，不会改变要返回的那个r值
//		r = r + 5
//	}(r)
//	return        //空的return
//}
