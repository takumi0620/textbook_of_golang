package main

import (
	"fmt"
	"strconv"
)

/**
型変換について
*/
func main() {
	list_10_6()
}

func list_10_6() {
	yesNo := "no"
	launch := (yesNo == "yes")
	fmt.Println("Ready for launch:", launch)
}

func list_10_5() {
	// ブールの変換
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}

	fmt.Println("Ready for launch:", yesNo)
}

func list_10_4() {
	// intから文字列
	countdown := 10
	str := "発射まで" + strconv.Itoa(countdown) + "秒"
	fmt.Println(str)
}

func list_10_3() {
	// runeとbyteをstringへ
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))
}

func list_10_2() {
	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(h)
}

func list_10_1() {
	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("私は火星では", marsAge, "歳です")
}
