package main

func lengthOfLongestSubstring(s string) int {
	longestLength := 0
	for p1 := 0; p1 < len(s); p1++ {
		visited := map[byte]bool{s[p1]: true}
		for p2 := p1 + 1; p2 < len(s); p2++ {
			curChar := s[p2]
			if _, ok := visited[curChar]; ok {
				break
			} else {
				visited[curChar] = true
			}
		}
		longestLength = max(longestLength, len(visited))
	}
	return longestLength
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
