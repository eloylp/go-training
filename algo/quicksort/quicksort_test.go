package quicksort_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/eloylp/go-training/algo/quicksort"
)

func Test_order(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"case 1", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quicksort.Do(tt.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchADoFunc(size int, sut func(a []int) []int) func(b *testing.B) {
	return func(b *testing.B) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		numbers := make([]int, size)
		for i := 0; i < size; i++ {
			numbers = append(numbers, r.Intn(10000))
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sut(numbers)
		}
	}
}

func BenchmarkDoFunctions(b *testing.B) {
	b.Run("Benchmark Do", BenchADoFunc(10, quicksort.Do))
	b.Run("Benchmark DoSimple", BenchADoFunc(10, quicksort.DoSimple))
}
