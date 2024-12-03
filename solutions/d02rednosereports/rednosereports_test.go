package d02rednosereports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckReports(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	resA, resB := CheckReports(input)
	assert.Equal(t, 2, resA)
	assert.Equal(t, 4, resB)
}
