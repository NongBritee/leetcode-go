package leetcode

func reverse(x int) int {

	sign := 1
	if x <= 0 {
		sign = -1
		x = x * -1 // abs
	}
	result := 0
	for x > 0 {
		n := x % 10
		result = (result * 10) + n
		x = x / 10
	}
	if result < -2147483648 || result > 2147483647 {
		return 0
	}
	return result * sign
}
