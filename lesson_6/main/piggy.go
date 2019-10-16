package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggy := 0.0

	for piggy < 20.00 {
		switch rand.Intn(3) {
		case 0:
			piggy += 0.05
		case 1:
			piggy += 0.10
		case 2:
			piggy += 0.25
		}
		fmt.Printf("現在の貯金額は%vドルです\n", piggy)
	}
}
