package main

import "fmt"

func main() {
	list_18_6()
}

func list_18_6() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)
}

// 可変長引数
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}

func list_18_5() {
	// makeでsliceを事前に定義
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs)
}

func list_18_4() {
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

	// インデックス3つでスライスの容量を制限できる かつ元の配列への参照を持たない
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "ケレス")
	worlds[1] = "ちきう"
	fmt.Println(planets)

	fmt.Println(worlds)
}

func list_18_3() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dwarfs = append(dwarfs, "Orcus")
	dump("dwarfs", dwarfs)
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs", dwarfs)
}

func list_18_2() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])
}

func dump(label string, slice []string) {
	fmt.Printf("%v: 長さ %v, 容量 %v %v\n", label, len(slice), cap(slice), slice)
}

func list_18_1() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
}
