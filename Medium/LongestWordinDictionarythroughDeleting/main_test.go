package main

import "testing"

func TestIsSubWord(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abcd", "a", true},
		{"a", "b", false},
		{"ab", "ba", false},
		{"a", "ab", false},
	}

	for _, test := range tests {
		if got := isSubWord(test.s, test.t); got != test.want {
			t.Errorf("string: %s, target: %s, want: %v, got: %v", test.s, test.t, test.want, got)
		}
	}
}

func TestFindLongestWord(t *testing.T) {
	tests := []struct {
		s    string
		arr  []string
		want string
	}{
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
	}

	for _, test := range tests {
		if got := findLongestWord(test.s, test.arr); got != test.want {
			t.Errorf("string: %s, arr: %v, want: %s, got: %s", test.s, test.arr, test.want, got)
		}
	}
}
