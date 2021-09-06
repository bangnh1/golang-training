package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		numbers  []int
		expected []int
	}{
		{[]int{1, 2, 5, 2, 6, 2, 5}, []int{5, 6, 1, 2}},
		{[]int{1, 2, 3, 4, 4, 7, 2, 5, 0, 1, 6}, []int{2, 3, 4, 7, 5, 0, 6, 1}},
		{[]int{11, 9, 1, 544, 4, 147, 200, 5, 20, 11, 5}, []int{4, 147, 200, 5, 20, 11, 9, 1, 544}},
	}
	for _, test := range tests {
		assert.ElementsMatch(removeDuplicates(test.numbers), test.expected)
	}
}
