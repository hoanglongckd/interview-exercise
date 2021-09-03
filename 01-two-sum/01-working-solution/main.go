package _1_working_solution

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		numberToFind := target - nums[i]

		for j := i + 1; j < len(nums); j++ {
			if nums[j] == numberToFind {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
