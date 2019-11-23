package main

import (
	"fmt"
	"image"
	"log"
	"sync"
	"time"
)

var mu sync.Mutex

// これまでに訪問したwebページを追跡管理する
type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

type RoverDriver struct {
	commandc chan command
}
type command int

const (
	right = command(0)
	left  = command(1)
)

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
				log.Printf("new direction %v", direction)
			}
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is", pos)
			next = time.After(time.Second)
		}
	}
}

// URLによるページ訪問回数を追跡管理し、更新したリンク数を返す
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}
