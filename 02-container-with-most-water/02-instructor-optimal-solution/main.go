package main

func findMaxArea(height []int) int {
	maxArea := 0

	if heightLens := len(height); heightLens > 1 {
		p1 := 0
		p2 := heightLens - 1

		for p1 < p2 {
			width := p2 - p1
			length := 0
			if height[p1] > height[p2] {
				length = height[p2]
				p2--
			} else {
				length = height[p1]
				p1++
			}
			curArea := width * length
			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}

	return maxArea
}
