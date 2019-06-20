package main

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"a.b,.", "a.b,."},
		{"", ""},
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
	}

	for _, test := range tests {
		if got := reverseVowels(test.s); got != test.want {
			t.Errorf("string: %s, want: %s, got: %s", test.s, test.want, got)
		}
	}
}
