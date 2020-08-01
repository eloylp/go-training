package climb

import (
	"testing"
)

func Test_climbStairs(t *testing.T) {

	tests := []struct {
		name  string
		steps int
		want  int
	}{
		{"For 1 step", 1, 1},
		{"For 2 step", 2, 2},
		{"For 3 step", 3, 3},
		{"For 4 step", 4, 5},
		{"For 5 step", 5, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cl(tt.steps); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
