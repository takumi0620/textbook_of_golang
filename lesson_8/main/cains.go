package main

import "fmt"

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	fmt.Println("大犬座矮小銀河と太陽の距離は光速で", distance/lightSpeed/secondsPerDay, "日かかる")
}
