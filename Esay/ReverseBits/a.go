package a

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	str := fmt.Sprintf("%b", num)
	str = Reverse(str)

	for len(str) < 32 {
		str += "0"
	}

	ans, err := strconv.ParseUint(str, 2, 32)
	if err != nil {
		return uint32(0)
	}

	return uint32(ans)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
