package a

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		x    uint32
		want uint32
	}{
		{43261596, 964176192},
	}

	for _, tt := range tests {
		if tt.want != reverseBits(tt.x) {
			t.Fail()
		}
	}
}
