package quicksort_test

import (
	"github.com/eloylp/go-training/algo/quicksort"
	"reflect"
	"testing"
)

func Test_order(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"case 1", []int{0, 8, 7, 6, 1, 4, 3, 9, 5, 2}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{"case 2", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quicksort.Do(tt.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("order() = %v, want %v", got, tt.want)
			}
		})
	}
}
