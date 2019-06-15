package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"", 0},
		{"bbbb", 1},
		{"pwwkew", 3},
	}

	for _, test := range tests {
		if got := lengthOfLongestSubstring(test.s); got != test.want {
			t.Errorf("input: %s, want: %d, got: %d", test.s, test.want, got)
		}
	}
}
