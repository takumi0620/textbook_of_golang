package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list_2_6()
}

func list_2_6() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)
}

func list_2_5() {
	// インクリメント色々 前置はサポートしない
	var age = 41
	age = age + 1
	age += 1
	age++

	fmt.Println(age)
}

func list_2_4() {
	var weight = 149.0
	// 代入 どっちも一緒
	weight = weight * 0.3784
	weight *= 0.3784

	fmt.Println(weight)
}

func list_2_3() {
	// 定数
	const lightSpeed = 299792
	// 変更可能
	var distance = 56000000
	fmt.Println(distance/lightSpeed, "秒")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "秒")

	// 纏めて宣言する方法いろいろ
	var a1 = 100
	var a2 = 200

	var (
		b1 = 100
		b2 = 200
	)

	var c1, c2 = 100, 200

	fmt.Println(a1, a2, b1, b2, c1, c2)
}

func list_2_2() {
	fmt.Printf("火星の表面で、私の体重は、%vポンド、", 149.0*0.3783)
	fmt.Printf("年齢は%v際になるでしょう\n", 41*365/687)

	// パディングの設定もできる 正だと左　負だと右
	fmt.Printf("パディング左4:%4v:end\n", "1")
	fmt.Printf("パディング右4:%-4v:end", "2")
}

func list_2_1() {
	fmt.Print("火星の表面で、私の体重は")
	fmt.Print(149.0 * 0.3783)
	fmt.Print("ポンド、年齢は")
	fmt.Print(41 * 365 / 687)
	fmt.Print("歳になるでしょう")
}
