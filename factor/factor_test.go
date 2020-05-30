package factor_test

import (
	"github.com/eloylp/go-training/factor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactor(t *testing.T) {
	samples := []struct{
		Context string
		Primes []int
		Number int
		Expected []int
	}{
		{"First case",[]int{2, 3, 5}, 30, []int{2, 3, 5}},
		{"Second case",[]int{2, 3, 5, 7}, 28, []int{2, 2, 7}},
	}
	for _, s := range samples {
		t.Run(s.Context, func(t *testing.T) {
			result := factor.Factor(s.Primes, s.Number)
			assert.Equal(t, s.Expected, result)
		})
	}
}
