package a

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	ans := 0
	sum := 0
	pos := 0
	rest := prices

	for {
		sum, pos, rest = findx(rest)
		ans += sum

		if pos < 0 || pos >= len(prices)-1 {
			break
		}
	}

	return ans
}

func findx(arr []int) (int, int, []int) {
	if len(arr) <= 1 {
		return 0, -1, []int{}
	}

	sum := 0
	i := 1
	pre := 0

	for ; i < len(arr); i++ {
		if arr[i] > arr[pre] {
			pre = i
			continue
		} else {
			sum = arr[i-1] - arr[0]
			return sum, i, arr[i:]
		}
	}

	return arr[len(arr)-1] - arr[0], -1, []int{}
}
