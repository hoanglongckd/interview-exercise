package _2_my_optimal_solution

func FirstTwoSum(nums []int, target int) []int {
	type numProp struct {
		exist bool
		index int
	}
	visitedNums := map[int]numProp{}

	for i := 0; i < len(nums); i++ {
		numberToFind := target - nums[i]
		currentMapVal := visitedNums[numberToFind]

		if currentMapVal.exist {
			return []int{currentMapVal.index, i}
		}
		visitedNums[nums[i]] = numProp{exist: true, index: i}
	}
	return []int{}
}

func SecondTwoSum(nums []int, target int) []int {
	type numProp struct {
		exist bool
		index int
	}
	visitedNums := map[int]numProp{}

	for i := 0; i < len(nums); i++ {
		currentMapVal := visitedNums[nums[i]]

		if currentMapVal.exist {
			return []int{currentMapVal.index, i}
		}
		numberToFind := target - nums[i]
		visitedNums[numberToFind] = numProp{exist: true, index: i}
	}
	return []int{}
}
