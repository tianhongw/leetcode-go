package main

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	tests := []struct {
		c    int
		want bool
	}{
		{0, true},
		{5, true},
		{3, false},
		{2, true},
	}

	for _, test := range tests {
		if got := judgeSquareSum(test.c); got != test.want {
			t.Errorf("int: %d, want: %v, got: %v", test.c, test.want, got)
		}
	}
}
