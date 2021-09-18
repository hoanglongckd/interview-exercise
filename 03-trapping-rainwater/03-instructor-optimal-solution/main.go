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
			if leftVal > leftMax {
				leftMax = leftVal
			} else {
				totalWater += leftMax - leftVal
			}
			left++
		} else {
			if rightVal > rightMax {
				rightMax = rightVal
			} else {
				totalWater += rightMax - rightVal
			}
			right--
		}
	}
	return totalWater
}
