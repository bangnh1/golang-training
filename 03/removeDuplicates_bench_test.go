package main

import "testing"

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeDuplicates([]int{1, 2, 5, 2, 6, 2, 5})
	}
}
