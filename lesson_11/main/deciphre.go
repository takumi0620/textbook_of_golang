package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	for i := 0; i < len(cipherText); i++ {
		value := cipherText[i] - (keyword[i%len(keyword)] - 'A')
		if value < 'A' {
			value = value + 26
		}
		fmt.Printf("%c", value)
	}
	fmt.Println("")
}
