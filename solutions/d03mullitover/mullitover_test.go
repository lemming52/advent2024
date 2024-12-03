package d03mullitover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuMulitply(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectedA int
		expectedB int
	}{
		{
			name:      "a",
			input:     "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expectedA: 161,
			expectedB: 161,
		}, {
			name:      "b",
			input:     "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expectedA: 161,
			expectedB: 48,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			resA, resB := SumMultiply(tt.input)
			assert.Equal(t, tt.expectedA, resA)
			assert.Equal(t, tt.expectedB, resB)
		})
	}

}
