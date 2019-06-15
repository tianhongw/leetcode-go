package main

func lengthOfLongestSubstring(s string) int {
	start := 0
	lastC := make(map[byte]int)
	ans := 0

	for i, c := range []byte(s) {
		lastI, ok := lastC[c]
		if lastI >= start && ok {
			start = lastI + 1
		}

		if i-start+1 > ans {
			ans = i - start + 1
		}

		lastC[c] = i
	}

	return ans
}
