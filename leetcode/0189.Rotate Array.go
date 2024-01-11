package leetcode

func rotate(nums []int, k int) {
	numLength := len(nums)
	if numLength <= 1 {
		return
	}
	k = k % numLength // when k >= len(nums)
	firstLen := numLength - k
	copy(nums, append(nums[firstLen:], nums[:firstLen]...))
}
