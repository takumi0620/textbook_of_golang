package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Universe [][]bool

const (
	width  = 80
	height = 15
)

func main() {
	universe := NewUniverse()
	universeII := NewUniverse()
	universe.Seed()

	universe.Show()

	for i := 0; i < 100; i++ {
		Step(universe, universeII)
		universe.Show()
		time.Sleep(time.Second / 30)
		universeII, universe = universe, universeII
	}
}

func Step(a, b Universe) {
	for y, row := range a {
		for x, _ := range row {
			b[y][x] = a.Next(x, y)
		}
	}
}

func (u Universe) Next(x, y int) bool {
	naighbors := u.Neighbors(x, y)
	return naighbors == 3 || naighbors == 2 && u.Alive(x, y)
}

/**
生きている隣接セルの数を返す
*/
func (u Universe) Neighbors(x, y int) int {
	aliveNum := 0
	// 左
	if u.Alive(x-1, y) {
		aliveNum++
	}

	// 左上
	if u.Alive(x-1, y-1) {
		aliveNum++
	}

	// 上
	if u.Alive(x, y-1) {
		aliveNum++
	}

	// 右上
	if u.Alive(x+1, y-1) {
		aliveNum++
	}

	// 右
	if u.Alive(x+1, y) {
		aliveNum++
	}

	// 右下
	if u.Alive(x+1, y+1) {
		aliveNum++
	}

	// 下
	if u.Alive(x, y+1) {
		aliveNum++
	}

	// 左下
	if u.Alive(x-1, y+1) {
		aliveNum++
	}
	return aliveNum
}

func (u Universe) Alive(x, y int) bool {
	// 範囲外の添字を指定された場合の考慮をする
	return u[(y+height)%height][(x+width)%width]
}

/**
マップの初期化
*/
func (u Universe) Seed() {
	for y, row := range u {
		for x, _ := range row {
			if rand.Intn(4)%3 == 0 {
				u[y][x] = true
			}
		}
	}
}

/**
マップを表示する
*/
func (u Universe) Show() {
	for _, row := range u {
		for _, cell := range row {
			if cell {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

/**
マップの作成
*/
func NewUniverse() Universe {
	universe := make([][]bool, height)
	for i, _ := range universe {
		universe[i] = make([]bool, width)
	}
	return universe
}
