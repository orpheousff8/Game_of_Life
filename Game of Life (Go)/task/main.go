package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	alive = 'O'
	dead  = ' '
)

type world [][]byte
type neighbour [][]byte

func (w world) print() {
	for i := range w {
		for j := range (w)[i] {
			fmt.Printf("%c", w[i][j])
		}
		fmt.Println()
	}
}

func (w world) clone() world {
	u := make(world, len(w))
	for i := range w {
		u[i] = make([]byte, len(w[i]))
		copy(u[i], w[i])
	}
	return u
}

func (w world) getNeighbour(I, J int) neighbour {
	size := len(w)
	n := make(neighbour, 3)

	for _, i := range [3]int{-1, 0, 1} {
		n[i+1] = make([]byte, 3)
		for _, j := range [3]int{-1, 0, 1} {
			n[i+1][j+1] = w[(I+i+size)%size][(J+j+size)%size]
		}
	}

	return n
}

func (w world) countAlive() int {
	count := 0
	for i := range w {
		for j := range (w)[i] {
			if w[i][j] == alive {
				count++
			}
		}
	}
	return count
}

func (n neighbour) print() {
	world(n).print()
}

func (n neighbour) isSurvive() bool {
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if n[i][j] == alive {
				count++
			}
		}
	}
	if n[1][1] == alive && (count == 2 || count == 3) {
		return true
	}
	if n[1][1] == dead && count == 3 {
		return true
	}
	return false
}

func newWorld(n int) world {
	var w = make(world, n)

	for i := range w {
		w[i] = make([]byte, n)
		for j := range w[i] {
			if rand.Intn(2) == 1 {
				w[i][j] = alive
			} else {
				w[i][j] = dead
			}
		}
	}

	return w
}

func main() {
	//initialise the world
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		return
	}

	mat1 := newWorld(n)

	//iterate through each generation
	for gen := 0; gen < 100; gen++ {
		//iterate into each cell
		mat2 := mat1.clone()

		for i := range mat2 {
			for j := range mat2[i] {
				x := mat2.getNeighbour(i, j)
				if x.isSurvive() {
					mat1[i][j] = alive
				} else {
					mat1[i][j] = dead
				}
			}
		}
		fmt.Printf("Generation #%d\n", gen+1)
		fmt.Printf("Alive: %d\n", mat1.countAlive())
		mat1.print()
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
	}
}
