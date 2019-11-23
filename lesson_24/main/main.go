package main

import (
	"fmt"
	"strings"
	"time"
)

var t interface {
	talk() string
}

type talker interface {
	talk() string
}

type stardater interface {
	YearDay() int
	Hour() int
}

type martian struct{}
type laser int
type starship struct {
	laser
}
type sol int
type location struct {
	lat, long float64
}

func main() {
	list_24_12()
}

func list_24_12() {
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func list_24_11() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))
	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", startdate(s))
}

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func startdate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

func list_24_8() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

}

func stardate(t time.Time) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

func list_24_7() {
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}

func list_24_6() {
	shout(martian{})
	shout(laser(2))
}

func list_24_5() {
	t = laser(3)
	shout(t)
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func list_24_3() {
	// ポリモフィズム
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())
}

func (m martian) talk() string {
	return "nack nack"
}

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}
