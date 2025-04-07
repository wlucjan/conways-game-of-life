package gameoflife

import "errors"

type Coordinates struct {
	X, Y int
}

type Pattern []Coordinates

type PatternBounds struct {
	MinX, MinY, MaxX, MaxY int
}

func NewPattern(c []Coordinates) (Pattern, error) {
	if len(c) == 0 {
		return nil, errors.New("pattern coordinates must not be empty")
	}

	return Pattern(c), nil
}

func (p Pattern) Center() Coordinates {
	bounds := p.bounds()
	width := bounds.MaxX
	height := bounds.MaxY

	centerX := width / 2
	centerY := height / 2

	return Coordinates{X: centerX, Y: centerY}
}

func (p Pattern) bounds() PatternBounds {
	bounds := PatternBounds{
		MinX: p[0].X, MinY: p[0].Y,
		MaxX: p[0].X, MaxY: p[0].Y,
	}

	for _, coord := range p[1:] {
		if coord.X < bounds.MinX {
			bounds.MinX = coord.X
		}
		if coord.X > bounds.MaxX {
			bounds.MaxX = coord.X
		}
		if coord.Y < bounds.MinY {
			bounds.MinY = coord.Y
		}
		if coord.Y > bounds.MaxY {
			bounds.MaxY = coord.Y
		}
	}
	return bounds
}
