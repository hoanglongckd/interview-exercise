package main

func findMaxArea(height []int) int {
	maxArea := 0
	p1 := 0
	p2 := len(height) - 1

	for p1 < p2 {
		width := p2 - p1
		length := min(height[p1], height[p2])
		maxArea = max(maxArea, width*length)

		if height[p1] > height[p2] {
			p2--
		} else {
			p1++
		}
	}
	return maxArea
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
