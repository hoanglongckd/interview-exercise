package _1_brute_force_solution

func trap(height []int) int {
	totalWater := 0
	for k, v := range height {
		maxLeft := 0
		maxRight := 0
		for left := k - 1; left >= 0; left-- {
			maxLeft = max(maxLeft, height[left])
		}
		for right := k + 1; right < len(height); right++ {
			maxRight = max(maxRight, height[right])
		}
		totalWater += positive(min(maxLeft, maxRight) - v)
	}
	return totalWater
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func positive(i int) int {
	if i > 0 {
		return i
	}
	return 0
}
