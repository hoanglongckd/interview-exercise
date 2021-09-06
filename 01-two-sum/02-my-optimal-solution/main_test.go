package _2_my_optimal_solution

import "testing"

var testCases = []struct {
	nums     []int
	target   int
	expected []int
}{
	{[]int{}, 1, []int{}},
	{[]int{1}, 1, []int{}},
	{[]int{1, 3}, 3, []int{}},
	{[]int{1, 3, 7, 9, 2}, 25, []int{}},
	{[]int{1, 5, 6}, 12, []int{}},
	{[]int{1, 6}, 7, []int{0, 1}},
	{[]int{1, 5, 6}, 7, []int{0, 2}},
	{[]int{1, 3, 7, 9, 2}, 11, []int{3, 4}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, item := range testCases {
		result := TwoSum(item.nums, item.target)
		if len(result) != len(item.expected) {
			t.Errorf("Want %v, Got %v", item.expected, result)
		} else if len(result) > 1 {
			if result[0] != item.expected[0] && result[1] != item.expected[1] {
				t.Errorf("Want %v, Got %v", item.expected, result)
			}
		}
	}
}
