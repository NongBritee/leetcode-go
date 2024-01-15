package leetcode

func moveZeroes(nums []int) {
	p1 := 0 // move when not found 0
	p2 := 0 // move when found 0
	n := len(nums)

	for p1 < n && p2 < n {
		if nums[p1] != 0 {
			p1++
		}
		if nums[p2] == 0 {
			p2++
		}

		if (p1 < n && p2 < n) && p2 > p1 {
			tmp := nums[p2]
			nums[p2] = nums[p1]
			nums[p1] = tmp
		}
		p2++

	}
}
