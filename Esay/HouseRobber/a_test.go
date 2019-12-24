package a

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
	}

	for _, tt := range tests {
		if tt.want != rob(tt.nums) {
			t.Fail()
		}
	}
}
