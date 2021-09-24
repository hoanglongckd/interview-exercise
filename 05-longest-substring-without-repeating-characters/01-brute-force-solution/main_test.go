package main

import "testing"

var testCases = []struct {
	s      string
	result int
}{
	{"abccabb", 3},
	{"cccccc", 1},
	{"", 0},
	{"abcbda", 4},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, testCase := range testCases {
		result := lengthOfLongestSubstring(testCase.s)
		if result != testCase.result {
			t.Errorf("With string: %s. Want %d, Got %d.", testCase.s, testCase.result, result)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdaeabcbdae")
	}
}