package leetcode

import "sort"

func removeElement(nums []int, val int) int {
	result := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 999
			result--
		}
	}
	sort.Ints(nums)

	return result
}
