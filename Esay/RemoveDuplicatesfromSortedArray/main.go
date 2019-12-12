package main

func removeDuplicates(nums []int) int {
	slowp := 0
	for fastp := 0; fastp < len(nums); fastp++ {
		if nums[fastp] != nums[slowp] {
			slowp++
			nums[slowp] = nums[fastp]
		}
	}

	return slowp + 1
}

func main() {

}
