package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"a", true},
		{"aba", true},
		{"abc", false},
		{"aabbc", false},
	}

	for _, test := range tests {
		if got := isPalindrome(test.s); got != test.want {
			t.Errorf("string: %s, want: %v, got: %v", test.s, test.want, got)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"a", "a"},
		{"aab", "aa"},
		{"babad", "aba"},
		{"babab", "babab"},
	}

	for _, test := range tests {
		if got := longestPalindrome(test.s); got != test.want {
			t.Errorf("string: %s, want: %s, got: %s", test.s, test.want, got)
		}
	}
}
