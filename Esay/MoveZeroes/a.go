package a

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
