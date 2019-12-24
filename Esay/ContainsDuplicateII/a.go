package a

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	gao := false

	for i, v := range nums {
		if pos, ok := m[v]; ok {
			if i-pos <= k {
				gao = true
			}
		}
		m[v] = i
	}

	return gao
}
