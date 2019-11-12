package main

import "fmt"

func main() {
	list_12_1()
}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func list_12_1() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "kは", celsius, "cです。")
}
