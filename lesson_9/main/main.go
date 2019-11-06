package main

import (
	"fmt"
	"unicode/utf8"
)

/**
多言語テキスト
*/
func main() {
	list_9_9()
}

func list_9_9() {
	question := "question of question"
	// rangeキーワードはコレクションの繰り返し
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}

func list_9_8() {
	question := "?como estas?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes\n", c, size)
}

func list_9_7() {
	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
}

func list_9_6() {
	// 1文字はシングルクォート
	c := 'a'
	// 文字列シフト
	c = c + 3
	fmt.Printf("%cは%[1]T型\n", c)
}

func list_9_5() {
	// 文字列インデックス
	// 文字列はイミュータブル
	message := "shalom"
	c := message[5]
	fmt.Printf("%c\n", c)
}

func list_9_4() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
}

func list_9_3() {
	// どっちもstring型
	fmt.Printf("%vは%[1]T型\n", "文字列")
	fmt.Printf("%vは%[1]T型", `文字列`)
}

func list_9_2() {
	// バッククォート使えば複数行も可能
	fmt.Println(`
1line
	2line
		3line 3line 3line
			4line`)
}

func list_9_1() {
	// 生の文字列
	fmt.Println("lorem ipsum\nlorem ipsum")
	// バッククォートでも可能
	fmt.Println(`outer haven`)
}
