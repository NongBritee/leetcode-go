package leetcode

import "fmt"

func FindLongestArrayOrder() {
	str := "15235137835692838387"
	asc := true
	tmpLongestArray := make([][]string, 0)
	maxSize := 0
	tmpArray := make([]string, 0)
	tmpArray = append(tmpArray, string(str[0]))

	for i := 1; i < len(str)-1; i++ {
		if asc { // asc
			if str[i] > str[i-1] {
				tmpArray = append(tmpArray, string(str[i]))
				if len(tmpArray) == maxSize {
					tmpLongestArray = append(tmpLongestArray, tmpArray)
				} else if len(tmpArray) > maxSize {
					tmpLongestArray = make([][]string, 0)
					tmpLongestArray = append(tmpLongestArray, tmpArray)
					maxSize = len(tmpArray)
				}
			} else {
				// reset tmpArray
				tmpArray = make([]string, 0)
				tmpArray = append(tmpArray, string(str[i]))
			}
		} else { // desc
			// ทำเหมือนกัน ปรับ > => <
		}
	}
	fmt.Printf("%v", tmpLongestArray)
}
