package insertionsort_test

import (
	"github.com/eloylp/go-training/algo/insertionsort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	sample := []int{56, 23, 78, 456, 45, 67}
	expected := []int{23, 45, 56, 67, 78, 456}
	insertionsort.InsertionSort(sample)
	assert.Equal(t, expected, sample)

}
