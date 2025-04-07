package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPattern_Center(t *testing.T) {
	testCases := []struct {
		name     string
		pattern  Pattern
		expected Coordinates
	}{
		{
			name:     "Single point pattern",
			pattern:  Pattern{{X: 5, Y: 5}},
			expected: Coordinates{X: 2, Y: 2},
		},
		{
			name:     "Block pattern",
			pattern:  Pattern{{3, 3}, {4, 3}, {3, 4}, {4, 4}},
			expected: Coordinates{X: 2, Y: 2},
		},
		{
			name:     "Glider pattern",
			pattern:  Pattern{{X: 1, Y: 0}, {X: 2, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}},
			expected: Coordinates{X: 1, Y: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.pattern.Center()

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNewPattern(t *testing.T) {
	testCases := []struct {
		name        string
		coords      []Coordinates
		expectError bool
	}{
		{
			name:        "Valid pattern",
			coords:      []Coordinates{{0, 0}, {1, 1}},
			expectError: false,
		},
		{
			name:        "Empty coordinates slice",
			coords:      []Coordinates{},
			expectError: true,
		},
		{
			name:        "Nil coordinates slice",
			coords:      nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPattern(tc.coords)

			if tc.expectError {
				if err == nil {
					t.Errorf("NewPattern() with input %v expected an error, but got nil", tc.coords)
				}
			} else {
				if err != nil {
					t.Errorf("NewPattern() with input %v expected no error, but got: %v", tc.coords, err)
				}
			}
		})
	}
}
