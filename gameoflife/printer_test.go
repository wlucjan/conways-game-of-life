package gameoflife_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlucjan/conways-game-of-life/gameoflife"
)

func TestPrinter(t *testing.T) {
	universe := gameoflife.NewUniverse(5, 5)
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
