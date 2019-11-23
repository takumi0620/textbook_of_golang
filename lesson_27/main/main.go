package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

var v interface{}

type number struct {
	value int
	valid bool
}

func main() {
	list_27_13()
}

func list_27_13() {
	n := newNumber(42)
	fmt.Println(n)
	e := number{}
	fmt.Println(e)
}

func (n number) String() string {
	if !n.valid {
		return "未設定"
	}
	return fmt.Sprintf("%d", n.value)
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func list_27_12() {
	fmt.Printf("%#v\n", v)
}

func list_27_11() {
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)
}

func list_27_10() {
	fmt.Printf("%T %v %v\n", v, v, v == nil)
}

func list_27_9() {
	var soup map[string]int
	fmt.Println(soup == nil)
	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)
	}
}

func list_27_7() {
	var soup []string
	fmt.Println(soup == nil)

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup))
	soup = append(soup, "onion", "carrot", "calery")
	fmt.Println(soup)
}

func list_27_6() {
	food := []string{"onion", "carrot", "calery"}
	sortStrings(food, nil)
	fmt.Println(food)
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return s[i] < s[j]
		}
	}
	sort.Slice(s, less)
}

func list_27_5() {
	var fn func(a, b int) int
	fmt.Println(fn == nil)
}

func list_27_3() {
	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}
