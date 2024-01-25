package main

import (
	"fmt"
	"myapp/04_package_example/mathutil"
)

func main() {
	// パッケージ内の関数の利用
	result := mathutil.Add(3, 5)
	fmt.Println("3 + 5 =", result)

	// パッケージ内の変数の利用
	fmt.Println("Pi value from mathutil package:", mathutil.Pi)
}
