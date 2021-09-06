package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax2Numbers(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		numbers  []int
		expected int
	}{
		{[]int{2, 1, 3, 4}, 3},
		{[]int{1, 2, 3, 4, 4, 7, 2, 5, 0, 1, 6}, 6},
		{[]int{11, 9, 1, 544, 4, 147, 200, 5, 20, 11, 5}, 200},
	}
	for _, test := range tests {
		assert.Equal(max2Numbers(test.numbers), test.expected)
	}
}
