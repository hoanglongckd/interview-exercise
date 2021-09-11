package main

func findMaxArea(height []int) int {
	maxArea := 0

	for p1 := 0; p1 < len(height)-1; p1++ {
		for p2 := p1 + 1; p2 < len(height); p2++ {
			p1Val := height[p1]
			p2Val := height[p2]

			if p1Val > p2Val {
				p1Val, p2Val = p2Val, p1Val
			}
			curArea := p1Val * (p2 - p1)

			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}
	return maxArea
}
