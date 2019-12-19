package a

func singleNumber(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}

	return ans
}
