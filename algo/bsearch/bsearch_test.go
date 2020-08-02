package bsearch_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/eloylp/go-training/algo/bsearch"
)

func TestLookUp(t *testing.T) {
	type args struct {
		s []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
		err  string
	}{
		{name: "Ordered 1", args: args{s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, v: 3}, want: 2},
		{name: "Not found search", args: args{s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, v: 12}, want: 0, err: "Not found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bsearch.LookUp(tt.args.s, tt.args.v)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
				return
			}
			assert.NoError(t, err)
			if got != tt.want {
				t.Errorf("LookUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createSoup(size int) ([]int, int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	soup := make([]int, size)
	for i := 0; i < size; i++ {
		soup[i] = r.Int()
	}
	worst := soup[size-1]
	sort.Ints(soup)
	return soup, worst
}

func BenchmarkLookUp(b *testing.B) {
	intSoup, search := createSoup(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = bsearch.LookUp(intSoup, search)
	}
}

func BenchmarkLinearLookUp(b *testing.B) {
	intSoup, search := createSoup(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = bsearch.LinearLookUp(intSoup, search)
	}
}
