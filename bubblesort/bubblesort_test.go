package bubblesort_test

import (
	"fmt"
	"github.com/eloylp/go-training/bubblesort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {

	sample := []int{45, 23, 89, 12, 34, 78}
	expected := []int{12, 23, 34, 45, 78, 89}
	bubblesort.BubbleSortInt(sample)
	assert.Equal(t, expected, sample)

}

func BenchmarkBubbleSortInt(b *testing.B) {
	for _, n := range []int{100, 400, 800, 1600 , 3200}{
		b.Run(fmt.Sprintf("Series of %v", n), func(b *testing.B) {
			sample := rand.Perm(n)
			b.ResetTimer()
			bubblesort.BubbleSortInt(sample)
		})
	}
}
