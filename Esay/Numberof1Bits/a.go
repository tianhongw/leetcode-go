package a

import "fmt"

func hammingWeight(num uint32) int {
	bstr := fmt.Sprintf("%b", num)
	c := 0

	for _, ch := range bstr {
		if ch == '1' {
			c++
		}
	}

	return c
}

func hammingWeight2(num uint32) int {
	c := 0
	for num > 0 {
		c++
		num = num & (num - 1)
	}

	return c
}
