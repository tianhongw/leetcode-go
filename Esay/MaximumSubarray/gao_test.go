package gao

import (
	"testing"

	ts "github.com/trytwice/leetcode-go/Util/Testing"
)

func TestGao(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, tt := range tests {
		ts.Equal(t, tt.want, maxSubArray(tt.arr))
	}
}
