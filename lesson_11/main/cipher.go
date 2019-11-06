package main

import "fmt"

func main() {
	cipher := "WEDIGYOULUVTHEGOPHERS"
	keyword := "GOLANG"

	for i := 0; i < len(cipher); i++ {
		value := cipher[i] + (keyword[i%len(keyword)] - 'A')
		if value > 'Z' {
			value = value - 26
		}
		fmt.Printf("%c", value)
	}
	fmt.Println("")
}
