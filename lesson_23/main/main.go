package main

import "fmt"

type sol int

type report struct {
	sol
	temperature
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	list_23_5()
}

func list_23_5() {
	report := report{sol: 15}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func list_23_4() {
	report := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: 78.0},
	}
	fmt.Printf("平均 %v C\n", report.average())
}

func list_23_3() {
	t := temperature{high: -1.0, low: -78.0}
	fmt.Printf("平均%v C\n", t.average())
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func list_23_2() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)

	fmt.Printf("おだやかな%v C\n", report.temperature.high)
}
