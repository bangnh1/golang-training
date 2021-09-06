package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxLengthElement(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		strings  []string
		expected []string
	}{
		{[]string{"aba", "aa", "ad", "c", "vcd"}, []string{"aba", "vcd"}},
		{[]string{"adfsdfba", "adssa", "afsdfsdd", "cfsdfdsfs", "fdsfssfsd"}, []string{"cfsdfdsfs", "fdsfssfsd"}},
	}

	for _, test := range tests {
		assert.Equal(findMaxLengthElement(test.strings), test.expected)
	}

}
