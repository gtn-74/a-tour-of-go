package main

import (
	"fmt"
)

func main() {
	// fmt.Println(math.E)
	// エクスポートされてない名前は、小文字でありアクセスできない様子
	// Ï fmt.Println(math.pi)
	// fmt.Println(math.Pi)
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}

//	func add(x int, y int) int {
//		return x + y
//	}
//
// 上を下のように型省略できる。
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
