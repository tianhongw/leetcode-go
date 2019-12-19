package a

import (
	"testing"

	ts "github.com/trytwice/leetcode-go/Util/Testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		ts.Equal(t, tt.want, maxProfit(tt.data))
	}
}
