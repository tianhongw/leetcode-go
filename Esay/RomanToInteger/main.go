package main

func romanToInt(s string) int {
	ans := 0
	l := len(s)

	for i := 0; i < l; i++ {
		switch s[i] {
		case 'M':
			ans += 1000
		case 'D':
			ans += 500
		case 'C':
			if i+1 < l && s[i+1] == 'D' {
				ans += 400
				i++
			} else if i+1 < l && s[i+1] == 'M' {
				ans += 900
				i++
			} else {
				ans += 100
			}
		case 'L':
			ans += 50
		case 'X':
			if i+1 < l && s[i+1] == 'L' {
				ans += 40
				i++
			} else if i+1 < l && s[i+1] == 'C' {
				ans += 90
				i++
			} else {
				ans += 10
			}
		case 'V':
			ans += 5
		case 'I':
			if i+1 < l && s[i+1] == 'V' {
				ans += 4
				i++
			} else if i+1 < l && s[i+1] == 'X' {
				ans += 9
				i++
			} else {
				ans++
			}
		}
	}

	return ans
}
