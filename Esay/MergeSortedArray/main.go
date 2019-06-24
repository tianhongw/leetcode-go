package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	mergeIndex := m + n - 1
	for index1 >= 0 || index2 >= 0 {
		if index1 < 0 {
			nums1[mergeIndex] = nums2[index2]
			index2--
		} else if index2 < 0 {
			nums1[mergeIndex] = nums1[index1]
			index1--
		} else if nums1[index1] >= nums2[index2] {
			nums1[mergeIndex] = nums1[index1]
			index1--
		} else {
			nums1[mergeIndex] = nums2[index2]
			index2--
		}
		mergeIndex--
	}
}
