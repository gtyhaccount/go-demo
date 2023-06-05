package main

//
//import "fmt"
//
//type sumable interface {
//	sum(int, int) int
//}
//
//// myFunc继承sumable接口
//type myFunc func(int) int
//
//func (f myFunc) sum (a, b int) int {
//	res := a + b
//	return f(res)
//}
//
//func sum10(num int) int {
//	return num * 10
//}
//
//func sum100(num int) int {
//	return num * 100
//}
//
//// icansum结构体继承sumable接口
//type icansum struct {
//	name string
//	res int
//}
//
//func (ics *icansum) sum(a, b int) int {
//	ics.res = a + b
//	return ics.res
//}
//
//func newSum(a int,b int) func(a,b int) (err error)  {
//	 c:= a + b
//	return func(a, b int) (err error) {
//		d:=c+a+b
//		fmt.Printf("%d",d)
//		return nil
//	}
//}
//
//
//// handler只要是继承了sumable接口的任何变量都行，我只需要你提供sum函数就好
//func handlerSum(handler sumable, a, b int) int {
//	res := handler.sum(a, b)
//	fmt.Println(res)
//	return res
//}
//
//func main() {
//	newFunc1 := myFunc(sum10)
//	newFunc2 := myFunc(sum100)
//
//	handlerSum(newFunc1, 2, 3) // 20
//	handlerSum(newFunc2, 1, 1) // 200
//
//	f:=newSum(1,2)
//	err := f(2, 2)
//	if err != nil {
//		println(err)
//		return
//	}
//	ics := &icansum{"I can sum", 0}
//	handlerSum(ics, 1, 1) // 2
//}
