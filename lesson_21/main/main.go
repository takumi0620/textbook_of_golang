package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var curiosity struct {
	lat float64
	lon float64
}

type location struct {
	lat, lon float64
}

type locationII struct {
	name     string
	lat, lon float64
}

func main() {
	list_21_10()
}

func list_21_10() {
	// JSON化するときの別名をつける
	type location struct {
		Lat float64 `json:"latitude"`
		Lon float64 `json:"longitude"`
	}

	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

func list_21_9() {
	// JSON化するにはフィールド名の先頭を大文字にする必要あり
	type location struct {
		Lat, Lon float64
	}

	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func list_21_8() {
	locations := []locationII{
		{name: "tanaka", lat: -4.5895, lon: 137.4417},
		{name: "hayashi", lat: -14.5684, lon: 175.472636},
		{name: "yamada", lat: -1.9462, lon: 354.4734},
	}

	fmt.Println(locations)
}

func list_21_7() {
	// 関連する値をただの変数で別々に定義するのはよくないという話
}

func list_21_6() {
	bradbury := location{-4.5895, 137.4417}
	curiosity := bradbury

	// ハードコピーされる
	curiosity.lon += 0.0106
	fmt.Println(bradbury, curiosity)
}

func list_21_5() {
	curiosity := location{-4.5895, 137.4417}

	// +をつけるとフィールド名も表示される。大きい構造体を調べる時に便利らしい
	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity)
}

func list_21_4() {
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)
}

func list_21_3() {
	opportunity := location{lat: -1.9462, lon: 354.4734}
	fmt.Println(opportunity)
	insight := location{lat: 4.5, lon: 135.9}
	fmt.Println(insight)
}

func list_21_2() {
	var split location
	split.lat = -14.5684
	split.lon = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.lon = 354.4734

	fmt.Println(split, opportunity)
}

func list_21_1() {
	curiosity.lat = -4.5895
	curiosity.lon = 137.4417
	fmt.Println(curiosity.lat, curiosity.lon)
	fmt.Println(curiosity)
}
