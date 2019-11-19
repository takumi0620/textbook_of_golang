package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	list_17_5()
}

func list_17_5() {
	planets := []string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"海王星",
		"冥王星",
	}

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}

func list_17_4() {
	planets := []string{" Venus ", "Earth ", " Mars"}
	hyperSpace(planets)
	fmt.Println(planets)
}

// 引数はslice
func hyperSpace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func list_17_3() {
	// 要素数省略するとsliceになる
	dwarfs := []string{
		"ケレス",
		"冥王星",
		"ハウメア",
		"マケマケ",
		"エリス",
	}

	fmt.Println(dwarfs)
}

func list_17_2() {
	planets := [...]string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"海王星",
		"冥王星",
	}

	// インデックスを省略すると0とlen(配列)になる
	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
}

func list_17_1() {
	planets := [...]string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"海王星",
		"冥王星",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants)
}
