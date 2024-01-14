package leetcode

import "testing"

func TestSingleNumberII(t *testing.T) {
	var cases = []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}

	for _, c := range cases {
		got := singleNumberII(c.nums)
		if got != c.want {
			t.Errorf("singleNumberII(%v) == %v, want %v", c.nums, got, c.want)
		}
	}
}
