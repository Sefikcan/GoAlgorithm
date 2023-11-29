package main

import "testing"

func BenchmarkSlice(b *testing.B) {
	n := 100

	b.Run("slice non alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := []int{}
			for i := 0; i < n; i++ {
				ints = append(ints, i)
			}
		}
	})

	b.Run("slice alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := make([]int, n)
			for i := 0; i < n; i++ {
				ints[i] = i
			}
		}
	})
}
