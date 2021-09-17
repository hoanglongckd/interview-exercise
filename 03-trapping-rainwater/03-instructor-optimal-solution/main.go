package _3_instructor_optimal_solution

func trap(height []int) int {
	totalWater := 0
	leftMax := 0
	rightMax := 0
	left := 0
	right := len(height) - 1

	for left < right {
		leftVal := height[left]
		rightVal := height[right]
		if leftVal <= rightVal {
			totalWater += positive(leftMax - leftVal)
			leftMax = max(leftMax, leftVal)
			left++
		} else {
			totalWater += positive(rightMax - rightVal)
			rightMax = max(rightMax, rightVal)
			right--
		}
	}
	return totalWater
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func positive(i int) int {
	if i > 0 {
		return i
	}
	return 0
}
