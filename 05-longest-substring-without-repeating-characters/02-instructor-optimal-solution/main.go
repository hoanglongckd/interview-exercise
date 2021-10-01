package main

func lengthOfLongestSubstring(s string) int {
	leftPtr := 0
	visitedChars := make(map[byte]int)
	longestSubStr := 0

	for rightPtr := 0; rightPtr < len(s); rightPtr++ {
		rightChar := s[rightPtr]

		if visitedIdx, ok := visitedChars[rightChar]; ok && leftPtr <= visitedIdx {
			leftPtr = visitedIdx + 1
		}
		longestSubStr = max(longestSubStr, rightPtr-leftPtr+1)
		visitedChars[rightChar] = rightPtr
	}
	return longestSubStr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
