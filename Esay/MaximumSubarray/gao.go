package gao

import "sort"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1]+nums[i], nums[i])
	}

	sort.Ints(dp)

	return dp[len(dp)-1]
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
