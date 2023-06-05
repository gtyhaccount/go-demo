package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	//计算随机数
	rand.Seed(15)
	fmt.Println(rand.Intn(10))
	//开方
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
	fmt.Println(math.Pi)
	fmt.Println(add(43, 32))
	fmt.Println(split(17))
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)

}
