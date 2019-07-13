package main

import (
	"testing"

	ts "github.com/trytwice/leetcode-go/Util/Testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		data []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, test := range tests {
		got := findKthLargest(test.data, test.k)
		ts.Equal(t, test.want, got)
	}
}
