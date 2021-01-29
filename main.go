package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

//cd c:\users\user\go\src\game_of_life

const (
	width  = 60
	height = 20
)

type Universe [][]bool

func NewUniverse() Universe {

	u := make(Universe, height)
	for i := 0; i < height; i++ {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for i := 0; i < height; i++ {
		s := ""
		for j := 0; j < width; j++ {
			if u[i][j] {
				s += "■"
			} else {
				s += " "
			}
		}
		fmt.Println(s)
	}
}

func (u Universe) Seed() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rand.Intn(4) == 3 {
				u[i][j] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	neighbors := u.Neighbors(x, y)
	return (neighbors == 3) || (neighbors == 2 && u.Alive(x, y))
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 600; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
	}
}
