package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "1",
			slice:    []int{2, 3, 9, 76, 5, 0, 0, 4},
			expected: []int{2, 3, 9, 76, 5, 4, 0, 0},
		},
		{
			name:     "2",
			slice:    []int{0, 0, 0, 1, 2, 3, 4},
			expected: []int{1, 2, 3, 4, 0, 0, 0},
		},
		{
			name:     "3",
			slice:    []int{8, 5, 9, 0, 0, 0},
			expected: []int{8, 5, 9, 0, 0, 0},
		},
		{
			name:     "4",
			slice:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "5",
			slice:    []int{7, 3, 0, 0, 0, 2, 4, 0, 5, 19},
			expected: []int{7, 3, 2, 4, 5, 19, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := run(test.slice)
			assert.Equal(t, test.expected, actual)

		})
	}
}
