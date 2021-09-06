package _2_my_optimal_solution

import "sort"

func TwoSum(nums []int, target int) []int {
	// [3, 2, 4]
	if len(nums) > 1 {
		sortedNums := nums    // [2, 3, 4] -- t = 6
		sort.Ints(sortedNums) // Take O(N*logN)
		p1 := 0
		p2 := len(sortedNums) - 1

		for p1 < p2 { // Take O(N)
			numberToFind := target - sortedNums[p1]

			if numberToFind == sortedNums[p2] {
				break
			}
			if sortedNums[p2] < numberToFind {
				p1++
				p2 = len(sortedNums) - 1
				continue
			}
			p2--
		}
		p1Value := sortedNums[p1]
		p2Value := sortedNums[p2]

		if p1Value != p2Value && p1Value+p2Value == target {
			var p1Index int
			var p2Index int

			for k, v := range nums { // Take O(N)
				if p1Value == v {
					p1Index = k
				}
				if p2Value == v {
					p2Index = k
				}
			}
			if p1Index > p2Index {
				p1Index, p2Index = p2Index, p1Index
			}
			return []int{p1Index, p2Index}
		}
	}
	// Total Time Complexity: O(N*logN) + O(N) + O(N) => O(N*(logN+3)) => O(N*logN)
	return []int{}
}
