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
