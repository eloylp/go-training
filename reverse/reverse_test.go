package reverse_test

import (
	"github.com/eloylp/go-training/reverse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	type sample struct {
		Input    []int
		Expected int
	}
	samples := []sample{
		{[]int{1, 1, 1, 1, 1}, 5},
		{[]int{1, 1, 1, 1, 0}, 4},
		{[]int{-1, 1, 1, 1, -1}, 1},
		{[]int{-1, 3, 1, 1, -1}, 3},
		{[]int{-1, -3, 1, 1, -1}, -3},
	}
	for _, s := range samples {
		assert.Equal(t, s.Expected, reverse.Sum(s.Input))
	}
}

func AssertReverseStringFunc(t *testing.T, f func(string) string) {
	type sample struct {
		Input    string
		Expected string
	}
	samples := []sample{
		{"eloy", "yole"},
		{"1234", "4321"},
	}
	for _, s := range samples {
		assert.Equal(t, s.Expected, f(s.Input))
	}
}

func TestReverseStringPreparedSlice(t *testing.T) {
	AssertReverseStringFunc(t, reverse.ReverseStringPreparedSlice)
}

func TestReverseStringIncrementalSlice(t *testing.T) {
	AssertReverseStringFunc(t, reverse.ReverseStringRune)
}

func BenchmarkReverseStringPreparedSlice(b *testing.B) {
	reverse.ReverseStringPreparedSlice("eloy")
}

func BenchmarkReverseStringIncrementalSlice(b *testing.B) {
	reverse.ReverseStringRune("eloy")
}
