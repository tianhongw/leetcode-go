package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	tmp := make([]string, 0)
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if isPalindrome(s[i:j]) {
				tmp = append(tmp, s[i:j])
			}
		}
	}

	ans := findMaxL(tmp)

	return ans
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func findMaxL(arr []string) string {
	p := 0
	l := len(arr[0])
	for i := 1; i < len(arr); i++ {
		if len(arr[i]) > l {
			l = len(arr[i])
			p = i
		}
	}

	return arr[p]
}
