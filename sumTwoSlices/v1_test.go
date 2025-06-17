package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumTwoSlices(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected []int
	}{
		{
			name:     "1",
			slice1:   []int{1, 2, 3},
			slice2:   []int{4, 5, 6},
			expected: []int{0, 5, 7, 9},
		},
		{
			name:     "2",
			slice1:   []int{5, 4, 4},
			slice2:   []int{4, 5, 6},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "3",
			slice1:   []int{9},
			slice2:   []int{9, 7, 5, 4, 3},
			expected: []int{0, 9, 7, 5, 5, 2},
		},
		{
			name:     "4",
			slice1:   []int{},
			slice2:   []int{4, 5, 6},
			expected: []int{0, 4, 5, 6},
		},
		{
			name:     "5",
			slice1:   []int{7, 8, 4, 8, 9},
			slice2:   []int{},
			expected: []int{0, 7, 8, 4, 8, 9},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getSumTwoSlices(test.slice1, test.slice2)
			assert.Equal(t, test.expected, actual)
		})
	}

}
