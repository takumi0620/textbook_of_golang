package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const marsDistance = 62100000

func main() {
	generateTitle(rand.Intn(20))
}

func generateTitle(ticketCount int) {
	fmt.Println("Spaceline          Days  Trip type    Price")
	fmt.Println("===========================================")
	for i := 0; i < ticketCount; i++ {
		tripTypeCode := getTripTypeCode()
		speed := getSpeed()
		charge := getCharge(speed, tripTypeCode)
		fmt.Println(getSpaceline(getSpaceLineCode()), getDays(speed), getTripType(tripTypeCode), "$"+strconv.Itoa(charge))
	}
}

func getDays(speed int) int {
	return marsDistance / (speed * 3600 * 24)
}

func getTripTypeCode() int {
	return rand.Intn(2)
}

func getSpaceLineCode() int {
	return rand.Intn(3)
}

func getSpeed() int {
	return rand.Intn(15) + 16
}

func getCharge(speed int, tripTypeCode int) int {
	return (speed + 20) * (tripTypeCode + 1)
}

func getSpaceline(spaceLineType int) string {
	ret := ""
	switch spaceLineType {
	case 0:
		ret = "Space Adventures  "
	case 1:
		ret = "SpaceX            "
	case 2:
		ret = "Virgin Galactic   "
	}
	return ret
}

func getTripType(tripType int) string {
	ret := ""
	switch tripType {
	case 0:
		ret = "Round-trip  "
	case 1:
		ret = "One-way     "
	}
	return ret
}
