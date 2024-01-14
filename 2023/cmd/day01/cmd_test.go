package day01

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected int64
		input    string
		fn       func(string) int64
	}{
		{
			expected: 142,
			input:    "test_1.txt",
			fn:       PartOne,
		},
		{
			expected: 281,
			input:    "test_2.txt",
			fn:       PartTwo,
		},
	}
	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}
