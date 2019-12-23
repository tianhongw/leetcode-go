package a

import (
	"math/rand"
	"testing"
	"time"
)

const (
	length = 1000000
)

var (
	data [length]uint32
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		data[i] = rand.Uint32()
	}
}

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		x    uint32
		want int
	}{
		{11, 3},
	}

	for _, tt := range tests {
		if tt.want != hammingWeight(tt.x) {
			t.Fail()
		}
	}
}

func BenchmarkHammingWeight1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hammingWeight(data[rand.Intn(length)])
	}
}

func BenchmarkHammingWeight2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hammingWeight2(data[rand.Intn(length)])
	}
}
