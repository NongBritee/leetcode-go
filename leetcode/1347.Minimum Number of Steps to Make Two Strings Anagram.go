package leetcode

func minSteps(s string, t string) int {
	sMap := make(map[rune]int, len(s))
	for _, ch := range s {
		sMap[ch] = sMap[ch] + 1
	}
	for _, ch := range t {
		if count, ok := sMap[ch]; ok && count > 0 {
			sMap[ch] = count - 1
		}
	}

	result := 0
	for _, v := range sMap {
		result = result + v
	}

	return result
}
