package main

import "fmt"

func main() {
	// スライスの宣言と初期化
	numbers := []int{1, 2, 3, 4, 5}

	// スライスへの要素の追加
	numbers = append(numbers, 6)

	// スライスの範囲指定
	slice := numbers[1:4]

	// マップの宣言と初期化
	person := map[string]string{
		"name": "John Doe",
		"age":  "30",
		"job":  "Engineer",
		"city": "New York",
	}

	// マップへの要素の追加
	person["email"] = "john@example.com"

	fmt.Println("Numbers:", numbers)
	fmt.Println("Slice:", slice)
	fmt.Println("Person:", person)
}
