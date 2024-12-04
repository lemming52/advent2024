package d04ceressearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSearch(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	resA, resB := WordSearch(input)
	assert.Equal(t, 18, resA)
	assert.Equal(t, 9, resB)
}
