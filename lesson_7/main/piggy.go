package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggy := 0.0
	for piggy < 2000 {
		switch rand.Intn(3) {
		case 0:
			piggy += 5
		case 1:
			piggy += 10
		case 2:
			piggy += 25
		}
		fmt.Println(piggy / 100)
	}
}
