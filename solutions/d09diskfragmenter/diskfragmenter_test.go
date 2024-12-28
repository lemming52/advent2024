package d09diskfragmenter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFragment(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "short",
			input:    "12345",
			expected: 60,
		}, {
			name:     "long",
			input:    "2333133121414131402",
			expected: 1928,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			resA := Fragment(tt.input)
			assert.Equal(t, tt.expected, resA)
		})
	}
}

func TestCompress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "short",
			input:    "12345",
			expected: 132,
		}, {
			name:     "long",
			input:    "2333133121414131402",
			expected: 2858,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			resA := Compress(tt.input)
			assert.Equal(t, tt.expected, resA)
		})
	}
}
