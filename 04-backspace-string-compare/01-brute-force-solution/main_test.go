package _1_brute_force_solution

import "testing"

var testCases = []struct {
	s, t   string
	result bool
}{
	{"", "", true},
	{"a#b#c#", "a#", true},
	{"abc##d", "ad", true},
	{"abc#d", "abd", true},
	{"aa#b", "Aa#b", false},
	{"a#b#", "", true},
	{"", "az##", true},
	{"abc#", "abc##", false},
	{"y#fo##f", "y#f#o##f", true},
}

func TestBackspaceCompare(t *testing.T) {
	for _, testCase := range testCases {
		result := backspaceCompare(testCase.s, testCase.t)
		if result != testCase.result {
			t.Errorf("Want %v, Got %v.", testCase.result, result)
		}
	}
}

func BenchmarkBackspaceCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompare(testCases[1].s, testCases[1].t)
	}
}
