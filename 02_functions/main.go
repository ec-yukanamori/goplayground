package main

import "fmt"

// 関数の宣言
func greet(name string) {
	fmt.Println("Hello, " + name + "!")
}

func add(a, b int) int {
	return a + b
}

func main() {
	// 関数の呼び出し
	greet("Alice")

	// 関数の戻り値の利用
	result := add(3, 7)
	fmt.Println("3 + 7 =", result)
}
