package leetcode

func removeDuplicates(nums []int) int {
	var result []int
	for i, n := range nums {
		if i == 0 || result[len(result)-1] != n {
			result = append(result, n)
		}
	}
	copy(nums, result)
	return len(result)
}
