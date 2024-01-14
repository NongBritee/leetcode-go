package leetcode

func singleNumberII(nums []int) int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n] = countMap[n] + 1
	}

	for k, v := range countMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
