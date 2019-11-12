package main

import "fmt"

type celsius float64
type kelvin float64

func main() {
	list_13_4(294.0)
}

func list_13_4(k kelvin) celsius {
	// リテラルは受け取れる
	return celsius(k - 273.15) // 型変換が必要
}

func list_13_3() {
	// 型を混ぜてつかうとエラーになる
	// type celsius float64
	// type fehrenheit float64
	// var c celsius = 20
	// var f fehrenheit = 20
	// if c == f {

	// }
}

func list_13_2() {
	type celsius float64
	const degrees = 20
	var temperature celsius = degrees
	fmt.Println(temperature)
}

func list_13_1() {
	// 独自の型を定義する
	type celsius float64

	var temperature celsius = 20
	fmt.Println(temperature)
}
