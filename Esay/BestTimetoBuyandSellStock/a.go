package a

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			ans = maxInt(ans, prices[i]-min)
		} else {
			min = prices[i]
		}
	}

	return ans
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
