package main

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", true},
		{"abbca", true},
		{"a", true},
		{"acbba", true},
		{"aba", true},
		{"abca", true},
		{"abcd", false},
		{"abcdefa", false},
		{"eeccccbebaeeabebccceea", false},
	}

	for _, test := range tests {
		if got := validPalindrome(test.s); got != test.want {
			t.Errorf("string: %s, want: %v, got: %v", test.s, test.want, got)
		}
	}
}
