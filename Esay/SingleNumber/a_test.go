package a

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tests {
		if tt.want != singleNumber(tt.data) {
			t.Fail()
		}
	}
}
