package main

func lengthOfLongestSubstring(s string) int {
	leftPtr := 0
	rightPtr := 0
	visitedChars := make(map[byte]int)
	longestSubStr := 0

	for rightPtr < len(s) {
		rightChar := s[rightPtr]

		if idx, ok := visitedChars[rightChar]; ok {
			if leftPtr > idx {
				longestSubStr = max(longestSubStr, rightPtr-leftPtr+1)
			} else {
				leftPtr = idx + 1
			}
		} else {
			longestSubStr = max(longestSubStr, rightPtr-leftPtr+1)
		}
		visitedChars[rightChar] = rightPtr
		rightPtr++
	}
	return longestSubStr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
