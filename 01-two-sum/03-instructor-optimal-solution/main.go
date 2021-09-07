package _2_my_optimal_solution

func FirstTwoSum(nums []int, target int) []int {
	visitedNums := map[int]int{}

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]
		numberToFind := target - currentNum

		if currentMapVal, ok := visitedNums[numberToFind]; ok {
			return []int{currentMapVal, i}
		}
		visitedNums[currentNum] = i
	}
	return []int{}
}

func SecondTwoSum(nums []int, target int) []int {
	visitedNums := map[int]int{}

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]

		if currentMapVal, ok := visitedNums[currentNum]; ok {
			return []int{currentMapVal, i}
		}
		numberToFind := target - nums[i]
		visitedNums[numberToFind] = i
	}
	return []int{}
}
