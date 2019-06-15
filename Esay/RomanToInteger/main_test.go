package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		if got := romanToInt(test.s); got != test.want {
			t.Errorf("string: %s, want: %d, got: %d", test.s, test.want, got)
		}
	}
}
