package main

import (
	"testing"

	ts "github.com/trytwice/leetcode-go/Util/Testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"", true},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tt := range tests {
		ts.Equal(t, tt.want, isValid(tt.str))
	}
}
