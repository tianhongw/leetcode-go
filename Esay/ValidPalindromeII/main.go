package main

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
	}
	return true
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
