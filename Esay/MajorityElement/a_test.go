package a

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tt := range tests {
		if tt.want != majorityElement(tt.data) {
			t.Fail()
		}
	}
}
