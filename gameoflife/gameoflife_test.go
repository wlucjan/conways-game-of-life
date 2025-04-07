package gameoflife_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlucjan/conways-game-of-life/gameoflife"
)

var printer = gameoflife.NewPrinter("O ", ". ")

func TestNewUniverse(t *testing.T) {
	newUniverse := gameoflife.NewUniverse(5, 5, nil)

	assert.Equal(t, 5, len(newUniverse))
	assert.Equal(t, 5, len(newUniverse[0]))
}

func TestNewUniverse_with_seed(t *testing.T) {
	seed, err := gameoflife.NewPattern([]gameoflife.Coordinates{{0, 0}})
	assert.NoError(t, err)

	newUniverse := gameoflife.NewUniverse(3, 3, seed)

	expected := ". . . \n" +
		". O . \n" +
		". . . "

	assert.Equal(t, expected, printer.Print(newUniverse))
}

func TestUniverse_Next(t *testing.T) {
	universe := gameoflife.NewUniverse(5, 5, gameoflife.GliderPattern)

	nextUniverse := universe.Next()

	expected := ". . . . . \n" +
		". . . . . \n" +
		". O . O . \n" +
		". . O O . \n" +
		". . O . . "

	assert.Equal(t, expected, printer.Print(nextUniverse))
}
