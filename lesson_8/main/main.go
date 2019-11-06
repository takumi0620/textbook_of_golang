package main

import (
	"fmt"
	"math/big"
)

/**
*ビッグナンバーだって
 */
func main() {
	list_8_3()
}

func list_8_3() {
	// 定数なら長大な数値でも大丈夫
	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const days = distance / lightSpeed / secondsPerDay
	fmt.Println("アンドロメダ銀河までの距離は", days, "km")
}

func quickCheck_8_2() {
	quickCheck1 := big.NewInt(86400)
	quickCheck2 := new(big.Int)
	quickCheck2.SetString("86400", 10)
	fmt.Println(quickCheck1, quickCheck2)
}

func list_8_2() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("アンドロメダ銀河までの距離は", distance, "km")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("光の速度で", days, "日かかる")
}

func list_8_1() {
	const lightSpeed = 299792   // 光速
	const secondsPerDay = 86400 // 1日の秒数

	var distance int64 = 41.3e12
	fmt.Println("アルファケンタウリまでの距離", distance)

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("光の速度で", days, "日")
}
