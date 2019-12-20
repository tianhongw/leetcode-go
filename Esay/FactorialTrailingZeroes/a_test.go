package a

import (
	"testing"
)

func TestTrailingZeros(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{20, 4},
	}

	for _, tt := range tests {
		if tt.want != trailingZeroes(tt.x) {
			t.Fail()
		}
	}
}
