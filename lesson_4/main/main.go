package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list_4_6()
}

var era = "AD"

func list_4_7() {
	// list_4_6のリファクタリング
	// 写経するに及ばない
}

func list_4_6() {
	// スコープの話
	year := 2018
	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}

func list_4_5() {
	// 4_4に同じ
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}

func list_4_4() {
	// この省略記法は気に入らない
	// が、スコープをifブロックの中に限定できるメリットがある
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Aventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}

func list_4_3() {
	// list_4_2の省略形
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
}

func list_4_2() {
	var count = 0
	// list4_1の省略形
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
}

func list_4_1() {
	var count = 0
	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	}

	// 同じ意味
	var a = 0
	b := 0
	fmt.Println(a, b)
}
