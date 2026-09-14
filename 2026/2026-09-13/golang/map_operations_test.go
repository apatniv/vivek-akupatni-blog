package main

import (
	"fmt"
	"testing"
)

var sizes = []int{1_000, 2_000, 4_000, 8_000, 16_000, 32_000, 64_000, 128_000}
var M int64 = (1 << 61) - 1

func InsertAndLookup(values []int64) int64 {
	dict := make(map[int64]int)
	for i, num := range values {
		dict[num] = i
	}
	var count int64
	for _, num := range values {
		if v, ok := dict[num]; ok {
			count += int64(v)
		}
	}
	return count
}

func InsertAndLookupPreAllocated(values []int64) int64 {
	dict := make(map[int64]int, len(values))
	for i, num := range values {
		dict[num] = i
	}
	var count int64
	for _, num := range values {
		if v, ok := dict[num]; ok {
			count += int64(v)
		}
	}
	return count
}

func BenchmarkInserts(b *testing.B) {
	for _, n := range sizes {
		b.Run(fmt.Sprintf("size=%v", n), func(b *testing.B) {
			values := make([]int64, n)
			for i := range n {
				values[i] = int64(i*1) * M
			}
			for b.Loop() {
				InsertAndLookup(values)
			}
		})

	}
}

func BenchmarkInsertsPreAllocated(b *testing.B) {
	for _, n := range sizes {
		b.Run(fmt.Sprintf("size=%v", n), func(b *testing.B) {
			values := make([]int64, n)
			for i := range n {
				values[i] = int64(i*1) * M
			}
			for b.Loop() {
				InsertAndLookupPreAllocated(values)
			}
		})

	}
}
