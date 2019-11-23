package main

import (
	"fmt"
	"strings"
)

type person struct {
	name, superpower string
	age              int
}

type stats struct {
	level             int
	endurance, health int
}

type character struct {
	name  string
	stats stats
}

type talker interface {
	talk() string
}

type martian struct{}

func main() {
	list_26_20()
}

type laser int

func list_26_20() {
	pew := laser(2)
	shout(&pew)
}

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func list_26_19() {
	shout(martian{})
	shout(&martian{})
}

func (m martian) talk() string {
	return "nack nack"
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func list_26_18() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Naptune", "Pluto",
	}

	reclassify(&planets)
	fmt.Println(planets)
}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func list_26_17() {
	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c", board[0][0])
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func list_26_16() {
	player := character{name: "Matthias"}
	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)

}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func list_26_13() {
	nathan := person{
		name: "Nathan",
		age:  17,
	}
	// ポインタレシーバじゃなくても大丈夫
	nathan.birthday()
	fmt.Printf("%+v\n", nathan)
}

func list_26_12() {
	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry)
}

func (p *person) birthday() {
	p.age++
}

func list_26_10() {
	redecca := person{
		name:       "Redecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&redecca)
	fmt.Printf("%+v\n", redecca)
}

func birthday(p *person) {
	p.age++
}

func list_26_8() {
	superpowers := &[3]string{"flight", "invisiblity", "super strength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
}

func list_26_7() {
	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	fmt.Println(timmy)
	fmt.Println(*timmy)

	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}

func list_26_5() {
	var administrator *string
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)
}

func list_26_4() {
	canada := "Canada"
	var home *string
	fmt.Printf("homeは%Tです\n", home)

	home = &canada
	fmt.Println(*home)
}

func list_26_3() {
	answer := 42
	address := &answer
	fmt.Printf("addressの型は%Tです\n", address)

	value := "value"
	fmt.Printf("valueの型は%Tです\n", &value)
}

func list_26_2() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)
}

func list_26_1() {
	answer := 42
	fmt.Println(&answer)
}
