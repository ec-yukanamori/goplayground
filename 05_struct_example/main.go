package main

import "fmt"

// Person 構造体の定義
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// 構造体の初期化
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// 構造体フィールドへのアクセス
	fmt.Println("First Name:", person1.FirstName)
	fmt.Println("Last Name:", person1.LastName)
	fmt.Println("Age:", person1.Age)

	// 構造体の変更
	person1.Age = 31

	// 新しい構造体の作成 (フィールドの一部のみ指定)
	person2 := Person{
		FirstName: "Alice",
		Age:       25,
	}

	fmt.Println("Updated Age:", person1.Age)
	fmt.Println("New Person:", person2)
}
