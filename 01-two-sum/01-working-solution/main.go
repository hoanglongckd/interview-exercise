package _1_working_solution

func TwoSum(nums []int, target int) []int {
	if len(nums) > 1 {
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
	}
	return []int{}
}
