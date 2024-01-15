package leetcode

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + 1
		digits[i] = sum
		if sum == 10 {
			digits[i] = 0
			continue
		}
		break
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
