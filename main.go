package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 20
	height = 20
)

type Universe [][]bool

func (u Universe) Seed() {
	var alive bool

	for ln, line := range u {
		for cn := range line {
			if rand.Intn(2) == 1 {
				alive = true
			} else {
				alive = false
			}

			u[ln][cn] = alive
		}
	}
}

func (u Universe) Show() {
	var pline [width]string

	for _, line := range u {
		for cn, alive := range line {
			if alive {
				pline[cn] = "1"
			} else {
				pline[cn] = "0"
			}
		}

		fmt.Println(pline)
	}
}

func (u Universe) Alive(x, y int) bool {
	return u[x][y]
}

func (u Universe) Neighbors(x, y int) int {
	var count int

	hasTopNeighbor := y > 0
	hasRightNeighbor := x < (width - 1)
	hasBottomNeighbor := y < (height - 1)
	hasLeftNeighbor := x > 0

	for _, neighbor := range [8]bool{
		hasTopNeighbor && u[y-1][x],
		hasRightNeighbor && u[y][x+1],
		hasBottomNeighbor && u[y+1][x],
		hasLeftNeighbor && u[y][x-1],
		hasTopNeighbor && hasRightNeighbor && u[y-1][x+1],
		hasBottomNeighbor && hasRightNeighbor && u[y+1][x+1],
		hasBottomNeighbor && hasLeftNeighbor && u[y+1][x-1],
		hasTopNeighbor && hasLeftNeighbor && u[y-1][x-1],
	} {
		if neighbor {
			count++
		}
	}

	return count
}

func (u Universe) Next(x, y int) bool {
	count := u.Neighbors(x, y)

	if u.Alive(x, y) {
		if count == 2 || count == 3 {
			return true
		}
	} else {
		if count == 3 {
			return true
		}
	}

	return false
}

func NewUniverse() Universe {
	universe := make(Universe, height)
	for ln := range universe {
		universe[ln] = make([]bool, width)
	}

	universe.Seed()

	return universe
}

func main() {
	universe_a := NewUniverse()
	universe_b := NewUniverse()

	universe_a.Show()

	for true {
		for ln, line := range universe_a {
			for cn := range line {
				universe_b[ln][cn] = universe_a.Next(ln, cn)
			}
		}

		universe_b.Show()
		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J")

		universe_a = universe_b
	}

}
