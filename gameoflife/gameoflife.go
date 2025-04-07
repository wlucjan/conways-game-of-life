package gameoflife

import "fmt"

const (
	alive = true
	dead  = false
)

type Universe [][]bool

func NewUniverse(width, height int, seed Pattern) Universe {
	if height <= 0 || width <= 0 {
		panic(fmt.Errorf("invalid universe dimensions: width=%d, height=%d must be positive", width, height))
	}

	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}

	if seed != nil {
		u.seed(seed)
	}

	return u
}

func (u Universe) Next() Universe {
	height := len(u)
	width := len(u[0])
	nextU := NewUniverse(width, height, nil)

	for y, row := range u {
		for x := range row {
			neighbors := u.countNeighbors(x, y)

			if u.isAlive(x, y) {
				if neighbors < 2 || neighbors > 3 {
					nextU[y][x] = dead // dies by underpopulation or overpopulation
				} else {
					nextU[y][x] = alive // survives
				}
			} else {
				if neighbors == 3 {
					nextU[y][x] = alive // becomes alive by reproduction
				} else {
					nextU[y][x] = dead // stays dead
				}
			}
		}
	}
	return nextU
}

func (u Universe) seed(p Pattern) {
	patternCenter := p.Center()
	universeCenter := u.center()

	for _, coord := range p {
		offsetX := coord.X - patternCenter.X
		offsetY := coord.Y - patternCenter.Y

		x := universeCenter.X + offsetX
		y := universeCenter.Y + offsetY

		if u.isWithinBounds(x, y) {
			u[y][x] = alive
		}
	}
}

func (u Universe) center() Coordinates {
	height := len(u)
	width := len(u[0])
	return Coordinates{X: width / 2, Y: height / 2}
}

func (u Universe) isWithinBounds(x, y int) bool {
	height := len(u)
	width := len(u[0])
	return y >= 0 && y < height && x >= 0 && x < width
}

func (u Universe) countNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			nx, ny := x+j, y+i

			if u.isWithinBounds(nx, ny) {
				if u.isAlive(nx, ny) {
					count++
				}
			}
		}
	}
	return count
}

func (u Universe) isAlive(x, y int) bool {
	if u.isWithinBounds(x, y) {
		return u[y][x] == alive
	}
	return false
}
