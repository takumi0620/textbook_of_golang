package main

/**
 * lesson7 整数
 */

import (
	"fmt"
	"time"
)

func main() {
	list_7_5()
}

func list_7_5() {
	// 120億超えても大丈夫
	future := time.Unix(12627780800, 0)
	fmt.Println(future)
}

func list_7_4() {
	// ビット桁溢れ
	var blue uint8 = 255
	fmt.Printf("%08b\n", blue)

	// += 10000とかはコンパイルエラー
	blue++
	fmt.Printf("%08b\n", blue)
}

func list_7_3() {
	// ビットを見る
	var green uint8 = 3
	fmt.Printf("%08b\n", green)

	green++
	fmt.Printf("%08b\n", green)
}

func list_7_2() {
	// 桁溢れで0に戻る
	var red uint8 = 255
	red++
	fmt.Println(red)

	// 符号付きだと-127
	var number int8 = 127
	number++
	fmt.Println(number)
}

func list_7_1() {
	year := 2018
	fmt.Printf("%T 型 %[1]v\n", year, year)

	days := 365.2425
	fmt.Printf("%T 型 %[1]v\n", days, days)
}
