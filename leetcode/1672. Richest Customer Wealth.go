package leetcode

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		sum := 0
		for _, bankWallet := range account {
			sum += bankWallet
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
