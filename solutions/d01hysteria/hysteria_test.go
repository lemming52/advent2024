package d01hysteria

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareIds(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	resA, resB := CompareIds(input)
	assert.Equal(t, 11, resA)
	assert.Equal(t, 31, resB)
}
