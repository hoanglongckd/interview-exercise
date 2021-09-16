package _2_my_optimal_solution

import "testing"

var testCases = []struct {
	height []int
	water  int
}{
	{
		[]int{},
		0,
	},
	{
		[]int{5},
		0,
	},
	{
		[]int{3, 4, 8},
		0,
	},
	{
		[]int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2},
		8,
	},
}

func TestTrap(t *testing.T) {
	for _, testCase := range testCases {
		water := trap(testCase.height)

		if water != testCase.water {
			t.Errorf("Want %d, Got %d.", testCase.water, water)
		}
	}
}

func BenchmarkTrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap(testCases[3].height)
	}
}
