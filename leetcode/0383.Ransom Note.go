package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	var mapMagazineCharWithCount map[string]int = make(map[string]int)
	for _, r := range magazine {
		char := string(r)
		mapMagazineCharWithCount[char] = mapMagazineCharWithCount[char] + 1
	}

	for _, r := range ransomNote {
		char := string(r)
		if mapMagazineCharWithCount[char] == 0 {
			return false
		} else {
			mapMagazineCharWithCount[char] = mapMagazineCharWithCount[char] - 1
		}
	}
	return true
}
