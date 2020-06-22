package popcount_test

import (
	"github.com/eloylp/go-training/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(uint64(i))
	}
}

func BenchmarkPopCountShiftValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShiftValue(uint64(i))
	}
}

func BenchmarkPopCountClearRightmost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClearRightmost(uint64(i))
	}
}
