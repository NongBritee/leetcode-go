package leetcode

func reverseString(s []byte) {

	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		tmp := s[i]
		s[i] = s[lastIndex]
		s[lastIndex] = tmp
		lastIndex--
	}

}
