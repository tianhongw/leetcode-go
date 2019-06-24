package main

func findLongestWord(s string, d []string) string {
	ans := ""

	for _, v := range d {
		if len(v) < len(ans) || (len(v) == len(ans) && v > ans) {
			continue
		}

		if isSubWord(s, v) {
			ans = v
		}
	}

	return ans
}

func isSubWord(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}

	return j == len(t)
}
