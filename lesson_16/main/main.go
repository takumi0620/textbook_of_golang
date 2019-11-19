package main

import "fmt"

func main() {
	list_16_8()
}

func list_16_8() {
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"

	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Println(board)
}

func list_16_7() {
	// 配列は値渡しだから関数に渡して中身更新しても何も意味がない、というLIST
}

func list_16_6() {
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

	// 配列はハードコピーされる
	planetsMarkII := planets
	planets[2] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}

func list_16_5() {
	dwarfs := [5]string{
		"ケレス",
		"冥王星",
		"ハウメア",
		"マケマケ",
		"エリス",
	}

	// rangeオブジェクトでループ
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}

func list_16_4() {
	dwarfs := [5]string{
		"ケレス",
		"冥王星",
		"ハウメア",
		"マケマケ",
		"エリス",
	}

	for i := 0; i < len(dwarfs); i++ {
		fmt.Println(i, dwarfs[i])
	}
}

func list_16_3() {
	// 要素数は初期化する時の要素の数になる
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

	fmt.Println(planets)
}

func list_16_2() {
	dwarfs := [5]string{
		"ケレス",
		"冥王星",
		"ハウメア",
		"マケマケ",
		"エリス",
	}

	fmt.Println(dwarfs)
}

func list_16_1() {
	var planets [8]string
	planets[0] = "水星"
	planets[1] = "金星"
	planets[2] = "地球"

	earth := planets[2]

	fmt.Println(earth)
}
