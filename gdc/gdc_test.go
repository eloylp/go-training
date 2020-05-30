package gdc_test

import (
	"fmt"
	"github.com/eloylp/go-training/gdc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGDC(t *testing.T) {

	samples := []struct {
		Input    []int
		Expected int
	}{
		{[]int{10, 5}, 5},
		{[]int{30, 15}, 15},
		{[]int{30, 9}, 3},
	}

	for _, s := range samples {
		t.Run(fmt.Sprintf("%v", s.Input), func(t *testing.T) {
			result := gdc.GDC(s.Input)
			assert.Equal(t, s.Expected, result)
		})
	}
}

func TestEuclideanGDC(t *testing.T) {

	samples := []struct {
		A    int
		B    int
		Expected int
	}{
		{10, 5, 5},
		{30, 15, 15},
		{30, 9, 3},
	}

	for _, s := range samples {
		t.Run(fmt.Sprintf("%v-%v", s.A, s.B), func(t *testing.T) {
			result := gdc.EuclideanGDC(s.A, s.B)
			assert.Equal(t, s.Expected, result)
		})
	}
}