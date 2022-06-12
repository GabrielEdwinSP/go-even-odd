package goevenodd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableOneParameter(t *testing.T) {
	tests := []struct {
		name     string
		request  int
		expected string
	}{
		{
			name:     "1",
			request:  1,
			expected: "Genap",
		},
		{
			name:     "2",
			request:  2,
			expected: "Genap",
		},
		{
			name:     "3",
			request:  3,
			expected: "Ganjil",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := EvenOdd(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestTableTwoParameter(t *testing.T) {
	tests := []struct {
		name     string
		request  []int
		expected []string
	}{
		{
			name:     "1",
			request:  []int{1, 2, 3},
			expected: []string{"Ganjil", "Genap", "Ganjil"},
		},
		{
			name:     "2",
			request:  []int{4, 2, 1},
			expected: []string{"Genap", "Genap", "Ganjil"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := EvenOddTwo(test.request...)
			assert.Equal(t, test.expected, result)
		})
	}
}
