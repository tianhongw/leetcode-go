package main

import "testing"

func TestRemoveDup(t *testing.T) {
	tests := []struct {
		intSlice []int
	}{
		{[]int{1, 1, 2}},
	}

	for _, tt := range tests {
		removeDuplicates(tt.intSlice)
	}
}
