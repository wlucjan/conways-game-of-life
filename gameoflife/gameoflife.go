package gameoflife

import "fmt"

type Universe [][]bool

func NewUniverse(width, height int) Universe {
	if height <= 0 || width <= 0 {
		panic(fmt.Errorf("invalid universe dimensions: width=%d, height=%d must be positive", width, height))
	}

	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}

	return u
}
