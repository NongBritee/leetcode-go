package leetcode

import "testing"

func TestPlusOneq(t *testing.T) {
	type testcases struct {
		name string
		in   []int
		out  []int
	}

	var testCases = []testcases{
		{"example1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"example2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"example3", []int{0}, []int{1}},
		{"example4", []int{9}, []int{1, 0}},
		{"example5", []int{9, 9}, []int{1, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := plusOne(tc.in)
			if len(res) != len(tc.out) {
				t.Errorf("expected %v, got %v", tc.out, res)
			}

			for i := range res {
				if res[i] != tc.out[i] {
					t.Errorf("expected %v, got %v", tc.out, res)
				}
			}
		})
	}

}
