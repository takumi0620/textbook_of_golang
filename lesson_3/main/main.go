package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	quick_check_3_2()
}

func quick_check_3_2() {
	const apple = "apple"
	const banana = "banana"
	const result = apple < banana
	fmt.Println(result)
}

func list_3_9() {
	var degrees = 0
	// 無限ループ whileは存在しない
	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

func list_3_8() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second) // 1秒sleep？
		count--
	}
	fmt.Println("Liftoff!")
}

func list_3_7() {
	var room = "lake"

	// if文みたいにも使える
	switch {
	case room == "cave":
		fmt.Println("きみは薄暗い洞窟のなかにいる")
	case room == "lake":
		fmt.Println("堅そうな氷が張っている")
		fallthrough // 条件の真偽に関係なく次のケース文を実行できる
	case room == "underwater":
		fmt.Println("水は氷るくらい冷たい")
	}
}

func list_3_6() {
	fmt.Println("洞窟の入り口だ。東へ進む道もある")
	var command = "go east"
	switch command {
	case "go east":
		fmt.Println("きみはさらに山を登る")
	case "enter cave", "go inside": // いずれか
		fmt.Println("きみは薄暗い洞窟の中にいる")
	case "read sign":
		fmt.Println("「未成年は立ち入り禁止」と書いてある")
	default:
		fmt.Println("なんだかよくわからない")
	}
}

func list_3_5() {
	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("ここは何も見えない")
	}
}

func list_3_4() {
	fmt.Println("2100年はうるう年ですか")
	var year = 2100
	// 論理演算子は||と&&
	// ||は真が見つかった以降は評価されない
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("はいうるう年です")
	} else {
		fmt.Println("違います")
	}
}

func list_3_3() {
	var command = "go insidea"

	// 条件に記述に()は不要
	if command == "go east" {
		fmt.Println("君はさらに山を登る")
	} else if command == "go inside" {
		fmt.Println("きみは洞窟に入り一生を過ごす")
	} else {
		fmt.Println("なんだかよくわからない")
	}
}

func list_3_2() {
	// 比較演算子は== < <= != > >=の6種類
	fmt.Println("入口の近くに「未成年立ち入り禁止」という立て札がある")

	var age = 41
	var minor = age < 18
	fmt.Printf("%v歳の僕は未成年か？ %v\n", age, minor)
}

func list_3_1() {
	fmt.Println("君は薄暗い洞窟の中にいる。")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("洞窟を出る", exit)
}
