package main

func reverseVowels(s string) string {
	buf := []byte(s)
	l, r := 0, len(buf)-1
	for l < r {
		if !isVowel(buf[l]) {
			l++
			continue
		}
		if !isVowel(buf[r]) {
			r--
			continue
		}
		if l < r {
			buf[l], buf[r] = buf[r], buf[l]
		}
		l++
		r--
	}

	return string(buf)
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}

	return false
}
