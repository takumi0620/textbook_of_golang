package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list_4_3()
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
