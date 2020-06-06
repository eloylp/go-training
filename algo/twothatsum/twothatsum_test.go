package twothatsum_test

import (
	"github.com/eloylp/go-training/algo/twothatsum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoThatSumFirstOccurrence(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9}
	resultA, resultB := twothatsum.TwoThatSumFirstOccurrence(list, 9)
	assert.Equal(t, 1, resultA)
	assert.Equal(t, 6, resultB)
}

func TestTwoThatSumLessDifference(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9}
	resultA, resultB := twothatsum.TwoThatSumLessDifference(list, 9)
	assert.Equal(t, 3, resultA)
	assert.Equal(t, 4, resultB)
}
