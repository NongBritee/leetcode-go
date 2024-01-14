package leetcode

import "fmt"

func rotateImage(matrix [][]int) {

	// swap i,j - j,i
	tmp := 0
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			tmp = matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = tmp
		}
		fmt.Println(matrix)
	}

	// swap horizantal
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			tmp := matrix[i][len(matrix)-1-j]
			matrix[i][len(matrix)-1-j] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}
