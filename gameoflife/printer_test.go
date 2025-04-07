package gameoflife_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlucjan/conways-game-of-life/gameoflife"
)

func TestPrinter_empty_universe(t *testing.T) {
	universe := gameoflife.NewUniverse(5, 5, nil)
	const (
		alive = "O"
		dead  = "."
	)
	printer := gameoflife.NewPrinter(alive, dead)

	expected := ".....\n" +
		".....\n" +
		".....\n" +
		".....\n" +
		"....."

	assert.Equal(t, expected, printer.Print(universe))
}

func TestPrinter_with_seed(t *testing.T) {
	seed, err := gameoflife.NewPattern([]gameoflife.Coordinates{{1, 0}, {2, 1}, {0, 2}, {1, 2}, {2, 2}})
	assert.NoError(t, err)

	newUniverse := gameoflife.NewUniverse(3, 3, seed)
	printer := gameoflife.NewPrinter("O ", ". ")

	expected := ". O . \n" +
		". . O \n" +
		"O O O "

	assert.Equal(t, expected, printer.Print(newUniverse))
}
