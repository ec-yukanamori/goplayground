package main

import "fmt"

func main() {
	// if文の例
	x := 10
	if x > 5 {
		fmt.Println("xは5より大きい")
	} else {
		fmt.Println("xは5以下")
	}

	// forループの例
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// switch文の例
	day := "Saturday"
	switch day {
	case "Monday":
		fmt.Println("月曜日です")
	case "Saturday", "Sunday":
		fmt.Println("週末です")
	default:
		fmt.Println("その他の日です")
	}
}
