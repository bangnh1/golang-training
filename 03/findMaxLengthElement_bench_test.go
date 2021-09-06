package main

import (
	"testing"
)

func BenchmarkFindMaxLengthElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findMaxLengthElement([]string{"aba", "aa", "ad", "c", "vcd"})
	}
}
