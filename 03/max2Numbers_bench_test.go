package main

import "testing"

func BenchmarkMax2Numbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = max2Numbers([]int{2, 1, 3, 4})
	}
}
