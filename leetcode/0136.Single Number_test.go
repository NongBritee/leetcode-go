package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tests {
		got := singleNumber(tt.nums)
		if got != tt.want {
			t.Errorf("singleNumber(%v) return %v, want %v", tt.nums, got, tt.want)
		}
	}
}
