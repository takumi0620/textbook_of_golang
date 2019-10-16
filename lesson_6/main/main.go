package main

import (
	"fmt"
	"math"
)

func main() {
	list_6_7()
}

func list_6_7() {
	celsius := 21.0 // 摂氏
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "F") // 掛け算先にすると精度ましになる 69.8
}

func list_6_6() {
	celsius := 21.0                         // 摂氏
	fmt.Print((celsius/5.0*9.0)+32, ".F\n") // 69.800000000000F
	fmt.Print((9.0/5.0*celsius)+32, ".F\n")
}

func list_6_5() {
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	// 浮動小数点は多言語同様2進数で表現するため精度が悪い
	// 高い完全性を要求される金額計算などには用いないこと 通称丸め誤差
	piggyBank := 0.1
	piggyBank += 0.2

	fmt.Println(piggyBank)
}

func list_6_4() {
	third := 1.0 / 3
	// ゼロ埋めは0を1つ置く
	fmt.Printf("%05.2f\n", third)
}

func list_6_3() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
}

func list_6_2() {
	// デフォのゼロ値で勝手に初期化される。price := 0.0と同じ
	var price float64
	fmt.Println(price)
}

func list_6_1() {
	// デフォが倍精度
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi64)
	fmt.Println(pi32)
}
