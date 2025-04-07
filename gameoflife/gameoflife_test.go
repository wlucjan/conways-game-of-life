package gameoflife_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlucjan/conways-game-of-life/gameoflife"
)

func TestNewUniverse(t *testing.T) {
	newUniverse := gameoflife.NewUniverse(5, 5, nil)

	assert.Equal(t, 5, len(newUniverse))
	assert.Equal(t, 5, len(newUniverse[0]))
}

func TestNewUniverse_with_seed(t *testing.T) {
	seed, err := gameoflife.NewPattern([]gameoflife.Coordinates{{0, 0}})
	assert.NoError(t, err)

	newUniverse := gameoflife.NewUniverse(3, 3, seed)
	printer := gameoflife.NewPrinter("O ", ". ")

	expected := ". . . \n" +
		". O . \n" +
		". . . "

	assert.Equal(t, expected, printer.Print(newUniverse))
}
