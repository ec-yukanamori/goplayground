package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	// 関数終了時にWaitGroupを通知
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	// 関数終了時にWaitGroupを通知
	defer wg.Done()

	for char := 'a'; char < 'e'; char++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("%c\n", char)
	}
}

func main() {
	// WaitGroupの初期化
	var wg sync.WaitGroup

	// 2つのゴルーチンを開始
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)

	// 両方のゴルーチンの終了を待つ
	wg.Wait()
}
