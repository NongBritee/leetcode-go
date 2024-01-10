package leetcode

import (
	"testing"
)

func TestRunningSum(t *testing.T) {
	type testcase struct {
		name   string
		input  []int
		expect []int
	}

	tests := []testcase{
		{"test 1", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"test 2", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{"test 3", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := runningSum(test.input)
			if len(got) != len(test.expect) {
				t.Errorf("runningSum(%v) = %v, want %v", test.input, got, test.expect)
			}
			for i := 0; i < len(got); i++ {
				if got[i] != test.expect[i] {
					t.Errorf("runningSum(%v) = %v, want %v", test.input, got, test.expect)
				}
			}
		})
	}
}
