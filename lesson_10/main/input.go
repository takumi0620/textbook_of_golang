package main

import "fmt"

func main() {
	fmt.Println(sToB2("no"))
}

func sToB1(str string) bool {
	return str == "true" || str == "yes" || str == "0"
}

func sToB2(str string) bool {
	return !(str == "false" || str == "no" || str == "0")
}
