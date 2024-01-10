package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	var result []string = make([]string, n)
	for i := 1; i <= n; i++ {
		idx := i - 1
		if i%3 == 0 {
			result[idx] += "Fizz"
		}
		if i%5 == 0 {
			result[idx] += "Buzz"
		}
		if result[idx] == "" {
			result[idx] = strconv.Itoa(i)
		}
	}
	return result
}
