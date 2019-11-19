package main

import (
	"fmt"
	"math"
)

func main() {
	list_19_5()
}

func list_19_5() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("セットのメンバ")
	}

	fmt.Println(set)
}

func list_19_4() {
	// スライスのマップ
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
}

func list_19_3() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2fの出現は%d回です\n", t, num)
	}
}

func list_19_2() {
	planets := map[string]string{
		"地球": "Sector ZZ9",
		"火星": "Sector ZZ9",
	}

	markII := planets
	planets["地球"] = "whoops"
	// 代入しても参照は同じ
	fmt.Println(planets)
	fmt.Println(markII)
}

func list_19_1() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temperature["Earth"])
	temperature["Earth"] = 16
	fmt.Println(temperature["Earth"])
}
