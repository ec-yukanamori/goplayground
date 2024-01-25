package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		// エラーが発生した場合、errors パッケージを使用してエラーメッセージを返す
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func main() {
	result, err := divide(8, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// エラーが発生した場合の処理
	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
