package leetcode

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}

	for _, tt := range tests {
		rotateImage(tt.matrix)
		if !reflect.DeepEqual(tt.matrix, tt.want) {
			t.Errorf("rotateImage(%v) return %v, want %v", tt.matrix, tt.matrix, tt.want)
		}
	}
}
