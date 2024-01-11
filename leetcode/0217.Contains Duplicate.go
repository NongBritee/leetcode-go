package leetcode

func containsDuplicate(nums []int) bool {
	check := make(map[int]bool, len(nums))
	for _, num := range nums {
		if check[num] {
			return true
		} else {
			check[num] = true
		}
	}
	return false
}
