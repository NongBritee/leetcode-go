package leetcode

func numberOfSteps(num int) int {
	result := num
	step := 0

	if result%2 == 1 {
		result = result - 1
		step++
	}

	for result != 0 {
		result = result / 2
		step++

		if result%2 == 1 {
			result = result - 1
			step++
		}
	}
	return step
}
