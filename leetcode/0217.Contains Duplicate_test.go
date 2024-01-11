package leetcode

import "testing"

func TestContainDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3}, false},
	}

	for _, tt := range tests {
		got := containsDuplicate(tt.nums)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%v\n", tt, got)
		}
	}
}
