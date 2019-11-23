package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

func main() {
	list_22_4()
}

func list_22_4() {
	mars := world{1.0}
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func list_22_3() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(newLocation(lat, long))
}

// コンストラクタっぽいものを自分で作る
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func list_22_2() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal())
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
