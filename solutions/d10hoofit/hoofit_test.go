package d10hoofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTrailHeads(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		expectedA int
		expectedB int
	}{
		{
			name: "1",
			input: []string{
				"..90..9",
				"...1.98",
				"...2..7",
				"6543456",
				"765.987",
				"876....",
				"987....",
			},
			expectedA: 4,
			expectedB: 13,
		}, {
			name: "2",
			input: []string{
				"10..9..",
				"2...8..",
				"3...7..",
				"4567654",
				"...8..3",
				"...9..2",
				".....01",
			},
			expectedA: 3,
			expectedB: 3,
		},
		{
			name: "1",
			input: []string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			expectedA: 36,
			expectedB: 81,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			resA, resB := FindTrailHeads(tt.input)
			assert.Equal(t, tt.expectedA, resA)
			assert.Equal(t, tt.expectedB, resB)
		})
	}
}
