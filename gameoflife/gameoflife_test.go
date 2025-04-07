package gameoflife_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlucjan/conways-game-of-life/gameoflife"
)

func TestNewUniverse(t *testing.T) {
	newUniverse := gameoflife.NewUniverse(5, 5)

	assert.Equal(t, 5, len(newUniverse))
	assert.Equal(t, 5, len(newUniverse[0]))
}
