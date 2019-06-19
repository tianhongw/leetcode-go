package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			numbers: []int{1, 2, 3, 5},
			target:  6,
			want:    []int{1, 4},
		},
	}

	for _, test := range tests {
		if got := twoSum(test.numbers, test.target); !equal(test.want, got) {
			t.Errorf("array: %v, target: %d, want: %v, got: %v", test.numbers, test.target, test.want, got)
		}
	}
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
