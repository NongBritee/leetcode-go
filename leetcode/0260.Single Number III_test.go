package leetcode

import (
	"sort"
	"testing"
)

func TestSingleNumberIII(t *testing.T) {
	type testscase struct {
		nums   []int
		expect []int
	}

	testcases := []testscase{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{-1, 0}, []int{-1, 0}},
		{[]int{0, 1}, []int{1, 0}},
	}

	for _, tc := range testcases {
		output := singleNumberIII(tc.nums)
		if len(output) != len(tc.expect) {
			t.Fatalf("input: %v, output: %v, expect: %v\n", tc.nums, output, tc.expect)
		}
		sort.Ints(output)
		sort.Ints(tc.expect)
		for i := 0; i < len(output); i++ {
			if output[i] != tc.expect[i] {
				t.Fatalf("input: %v, output: %v, expect: %v\n", tc.nums, output, tc.expect)
			}
		}
	}
}
