package _2_my_optimal_solution

func trap(height []int) int {
	rightHeight := getRightHeight(height)
	leftMax := 0
	totalWater := 0

	for idx, val := range height {
		rightMax := rightHeight[idx+1]
		minHeight := min(leftMax, rightMax)
		totalWater += positive(minHeight - val)
		if val > leftMax {
			leftMax = val
		}
	}
	return totalWater
}

func getRightHeight(height []int) map[int]int {
	rightHeight := map[int]int{}
	rightHeight[len(height)] = 0

	for right := len(height) - 1; right > 0; right-- {
		rightVal := height[right]
		rightMax := rightHeight[right+1]
		rightHeight[right] = max(rightVal, rightMax)
	}
	return rightHeight
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
