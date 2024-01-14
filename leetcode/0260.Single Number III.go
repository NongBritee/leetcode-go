package leetcode

func singleNumberIII(nums []int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n] = countMap[n] + 1
	}

	var result []int
	for k, v := range countMap {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}
