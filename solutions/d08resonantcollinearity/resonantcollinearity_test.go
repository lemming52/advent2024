package d08resonantcollinearity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAntiNodes(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
	resA, resB := FindAntiNodes(input)
	assert.Equal(t, 14, resA)
	assert.Equal(t, 34, resB)
}
